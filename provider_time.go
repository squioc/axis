package axis

import "time"

// Time is the time-based provider
type Time struct {
}

// Current gets the number of milliseconds elapsed since 1, January 1970
func (t *Time) Current() Position {
	return Position(time.Now().UnixNano())
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

// AfterFunc waits for the given distance to elapse
// And then calls the callback 
func (t *Time) AfterFunc(distance Distance, callback func()) Watcher {
	return NewTimeWatcher(time.AfterFunc(time.Duration(distance), callback))
}

// AfterChan waits for the given distance to elapse
// And then sends the new time on the given channel
func (t *Time) AfterChan(distance Distance, channel chan Position) Watcher {
	return t.AfterFunc(distance, func() {
		channel <- t.Current()
	})
}

// Since return the distance traveled since position
func (t *Time) Since(position Position) Distance {
	return Distance(t.Current() - position)
}
