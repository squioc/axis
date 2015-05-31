package axis

import "time"

// TimeWatcher wraps a time.Timer
type TimeWatcher struct {
	Timer *time.Timer
}

// NewTimeWatcher creates a new time-based watcher from a time.Timer
func NewTimeWatcher(timer *time.Timer) *TimeWatcher {
	return &TimeWatcher{Timer: timer}
}

// Reset changes the watcher to fire after the given distance
// Return True if the watcher changed, False otherwise
func (s *TimeWatcher) Reset(distance Distance) bool {
	return s.Timer.Reset(time.Duration(distance))
}

// Stop prevents the watcher from firing
// Return True if the watcher stopped, False if the watcher already fired or stopped
func (s *TimeWatcher) Stop() bool {
	return s.Timer.Stop()
}
