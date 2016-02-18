package axis

// Positionable is the interface for positionable
// items on a axis
type Positionable interface {
	Current() Position
	Since(Position) Distance
}

// Sleepable is the interface for sleepable provider
type Sleepable interface {
	Sleep(Distance)
}

// Trigger is the interface that wraps methods
// to define triggers
type Trigger interface {
	After(Distance) <-chan Position
	AfterFunc(Distance, func(Position)) Watcher
	AfterChan(Distance, chan Position) Watcher
}

// Provider is the interface that wraps methods
// to manipulate position
type Provider interface {
        Positionable
        Sleepable
        Trigger
}

// UpdatableProvider is the interface which allow
// to update the position of the provider
type UpdatableProvider interface {
	Provider
	Update(Position)
}
