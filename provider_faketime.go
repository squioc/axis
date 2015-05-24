package axis

import (
    "sync"
)

type Timer struct {
    position Position
    distance Distance
    channel chan Position
}

type Timers []Timer

type FakeTime struct {
    position Position
    timers map[Position][] chan Position
    mu sync.Mutex
}

func NewFakeTime(position Position) *FakeTime {
    return &FakeTime{
        position: position,
        timers: make(map[Position][] chan Position),
    }
}

func (f *FakeTime) Current() Position {
    return f.position
}

func (f *FakeTime) Sleep(distance Distance) {
    f.Update(addDistance(f.position, distance))
}

func (f *FakeTime) After(distance Distance) <-chan Position {
    c := make(chan Position, 1)
    f.AfterChan(distance, c)
    return c
}

func (f *FakeTime) AfterChan(distance Distance, channel chan Position) *FakeTimeWatcher {
    f.mu.Lock()
    defer f.mu.Unlock()

    until := addDistance(f.position, distance)
    f.timers[until] = append(f.timers[until], channel)
    return &FakeTimeWatcher{can_reset: true, can_stop: true}
}

func (f *FakeTime) Update(position Position) {
    f.position = position

    f.mu.Lock()
    defer f.mu.Unlock()
    for k, v := range f.timers {
        if k < f.position {
            for _, c := range v {
                c <- f.position
            }
            delete(f.timers, k)
        }
    }
}

func addDistance(position Position, distance Distance) Position {
    return Position(int64(position) + int64(distance))
}
