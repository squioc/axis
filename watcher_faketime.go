package axis

// FakeTimeWatcher simulates a watcher
type FakeTimeWatcher struct {
	canReset bool
	canStop  bool
}

// NewFakeTimeWatcher creates a new fake watcher
func NewFakeTimeWatcher(canReset, canStop bool) *FakeTimeWatcher {
	return &FakeTimeWatcher{canReset: canReset, canStop: canStop}
}

// Reset changes the watcher to fire after the given distance
// Return True if the watcher changed, False otherwise
func (s *FakeTimeWatcher) Reset(distance Distance) bool {
	return s.canReset
}

// Stop prevents the watcher from firing
// Return True if the watcher stopped, False if the watcher already fired or stopped
func (s *FakeTimeWatcher) Stop() bool {
	return s.canStop
}
