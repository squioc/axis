package axis


type FakeTimeWatcher struct {
    can_reset bool
    can_stop bool
}

func NewFakeTimeWatcher(can_reset, can_stop bool) *FakeTimeWatcher {
    return &FakeTimeWatcher{can_reset: can_reset, can_stop: can_stop}
}

func (s *FakeTimeWatcher) Reset(distance Distance) bool {
    return s.can_reset
}

func (s *FakeTimeWatcher) Stop() bool {
    return s.can_stop
}

