package axis

import "time"

// Time is the time-based provider
type Time struct {
}

// Current get the number of seconds elapsed since 1, January 1970
func (t *Time) Current() Position {
	return Position(time.Now().Unix())
}

// Sleep pauses the current goroutine for the given distance
func (t *Time) Sleep(distance Distance) {
	time.Sleep(time.Duration(distance))
}

// After waits for the given distance to elapse
// And then sends the new time on the returned channel
func (t *Time) After(distance Distance) <-chan Position {
	c := make(chan Position, 1)
	t.AfterChan(distance, c)
	return c
}

// AfterChan waits for the given distance to elapse
// And then sends the new time on the given channel
func (t *Time) AfterChan(distance Distance, channel chan Position) *TimeWatcher {
	f := func() {
		channel <- t.Current()
	}
	return NewTimeWatcher(time.AfterFunc(time.Duration(distance), f))
}
