package axis

type Provider interface {
    Current() Position
    Sleep(Distance)
    After(Distance) <-chan Position
    AfterChan(Distance, chan Position) *Watcher
}

type UpdatableProvider interface {
    Provider
    Update(Position)
}

