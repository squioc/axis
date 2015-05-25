package axis

import (
    "sync"
)

// FakeTime simulate a time-based provider
type FakeTime struct {
    // position gives the position of the faketime provider
    position Position

    // timers represents specific points in the time
    // it throws their position on a channel when it reaches them
    timers map[Position][] chan Position
    mu sync.Mutex
}

// Create a new faketime provider
func NewFakeTime(position Position) *FakeTime {
    return &FakeTime{
        position: position,
        timers: make(map[Position][] chan Position),
    }
}

// Get the current position of the provider
func (f *FakeTime) Current() Position {
    return f.position
}

// Pause the provider for the given distance
func (f *FakeTime) Sleep(distance Distance) {
    f.Update(addDistance(f.position, distance))
}

// Simulate a wait for the given distance to elapse
// and then sends the new position on the returned channel
func (f *FakeTime) After(distance Distance) <-chan Position {
    c := make(chan Position, 1)
    f.AfterChan(distance, c)
    return c
}

// Simulate a wait for the given distance to elapse
// and then sends the new position on the given channel
func (f *FakeTime) AfterChan(distance Distance, channel chan Position) *FakeTimeWatcher {
    f.mu.Lock()
    defer f.mu.Unlock()

    until := addDistance(f.position, distance)
    f.timers[until] = append(f.timers[until], channel)
    return &FakeTimeWatcher{can_reset: true, can_stop: true}
}

// Update the current position of the provider
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
