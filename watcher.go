package axis

// Watcher is the interface that allows
// to stop or reset an event associated to a point
type Watcher interface {
	// Reset the watcher
	Reset(Distance) bool

	// Stop the watcher
	Stop() bool
}
