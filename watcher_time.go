package axis

import "time"

type TimeWatcher struct {
    Timer *time.Timer
}

func NewTimeWatcher(timer *time.Timer) *TimeWatcher {
    return &TimeWatcher{Timer: timer }
}

func (s *TimeWatcher) Reset(distance Distance) bool {
    return s.Timer.Reset(time.Duration(distance))
}

func (s *TimeWatcher) Stop() bool {
    return s.Timer.Stop()
}

