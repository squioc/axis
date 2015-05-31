package axis

import (
	"testing"
)

func TestFakeTimeWatcherReset(t *testing.T) {
	// Arrange
	canReset := true
	canStop := true
	distance := Distance(0)
	watcher := NewFakeTimeWatcher(canReset, canStop)
	var actual bool

	// Act
	actual = watcher.Reset(distance)

	// Assert
	if actual != canReset {
		t.Fatalf("Mismatching expected boolean for the reset of the fake time watcher")
	}
}

func TestFaketimeWatcherStop(t *testing.T) {
	// Arrange
	canReset := true
	canStop := true
	watcher := NewFakeTimeWatcher(canReset, canStop)
	var actual bool

	// Act
	actual = watcher.Stop()

	// Assert
	if actual != canStop {
		t.Fatalf("Mismatching expected boolean for the stop of the fake time watcher")
	}
}
