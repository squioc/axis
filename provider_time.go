package axis

import "time"

type Time struct {
}

func (t *Time) Current() Position {
    return Position(time.Now().Unix())
}

func (t *Time) Sleep(distance Distance) {
    time.Sleep(time.Duration(distance))
}

func (t *Time) After(distance Distance) <-chan Position {
    c := make(chan Position, 1)
    t.AfterChan(distance, c)
    return c
}

func (t *Time) AfterChan(distance Distance, channel chan Position) *TimeWatcher {
    f := func() {
        channel <- t.Current()
    }
    return newTimeWatcher(time.AfterFunc(time.Duration(distance), f))
}
