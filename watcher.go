package axis

type Watcher interface {
    Reset(Distance) bool
    Stop() bool
}

