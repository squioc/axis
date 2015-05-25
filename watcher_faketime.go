package axis

// FakeTimeWatcher simulate a watcher
type FakeTimeWatcher struct {
    can_reset bool
    can_stop bool
}

// NewFakeTimeWatcher create a new fake watcher
func NewFakeTimeWatcher(can_reset, can_stop bool) *FakeTimeWatcher {
    return &FakeTimeWatcher{can_reset: can_reset, can_stop: can_stop}
}

// Reset change the watcher to fire after the given distance
// Return True if the watcher changed, False otherwise
func (s *FakeTimeWatcher) Reset(distance Distance) bool {
    return s.can_reset
}

// Stop prevent the watcher from firing
// Return True if the watcher stopped, False if the watcher already fired or stopped
func (s *FakeTimeWatcher) Stop() bool {
    return s.can_stop
}

