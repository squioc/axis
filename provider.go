package axis

// Provider is the interface that wraps methods
// to manipulate position
type Provider interface {
	Current() Position
	Sleep(Distance)
	After(Distance) <-chan Position
	AfterChan(Distance, chan Position) Watcher
}

// UpdatableProvider is the interface which allow
// to update the position of the provider
type UpdatableProvider interface {
	Provider
	Update(Position)
}
