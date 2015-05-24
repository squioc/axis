package axis

import (
    "testing"
)


func TestFakeTimeWatcherReset(t *testing.T) {
    // Arrange
    can_reset := true
    can_stop := true
    distance := Distance(0)
    watcher := NewFakeTimeWatcher(can_reset, can_stop)
    var actual bool

    // Act
    actual = watcher.Reset(distance)

    // Assert
    if actual != can_reset  {
        t.Fatalf("Mismatching expected boolean for the reset of the fake time watcher")
    }
}

func TestFaketimeWatcherStop(t *testing.T) {
    // Arrange
    can_reset := true
    can_stop := true
    watcher := NewFakeTimeWatcher(can_reset, can_stop)
    var actual bool

    // Act
    actual = watcher.Stop()

    // Assert
    if actual != can_stop  {
        t.Fatalf("Mismatching expected boolean for the stop of the fake time watcher")
    }
}
