package axis

import (
	"testing"
	"time"
)

func TestFakeTimeCurrent(t *testing.T) {
	// Arrange
	position := Position(1000)
	provider := NewFakeTime(position)
	var actual Position

	// Act
	actual = provider.Current()

	// Assert
	if actual != position {
		t.Fatalf("The current position mismatches the initial position")
	}
}

func TestFakeTimeUpdate(t *testing.T) {
	// Arrange
	position := Position(1000)
	newPosition := Position(2000)
	provider := NewFakeTime(position)
	var actual Position

	// Act
	provider.Update(newPosition)
	actual = provider.Current()

	// Assert
	if actual != newPosition {
		t.Fatalf("The current position mismatches the new position")
	}
}

func TestFakeTimeSleep(t *testing.T) {
	// Arrange
	position := Position(1000)
	distance := Distance(100)
	provider := NewFakeTime(position)
	expected := Position(int64(position) + int64(distance))
	var actual Position

	// Act
	provider.Sleep(distance)
	actual = provider.Current()

	// Assert
	if actual != expected {
		t.Fatalf("The actual position mismatches the expected position")
	}
}

func TestFakeTimeAfter(t *testing.T) {
	// Arrange
	position := Position(1000)
	distance := Distance(100)
	newPosition := Position(2000)
	provider := NewFakeTime(position)
	var actual Position

	// Act
	c := provider.After(distance)
	provider.Update(newPosition)

	// Assert
	select {
	case actual = <-c:
		if actual < newPosition {
			t.Fatalf("The actual position was fired too early")
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("Timeout. the test exceed the expected duration")
	}
}

func TestFakeTimeAfterFunc(t *testing.T) {
	// Arrange
	position := Position(1000)
	distance := Distance(100)
	newPosition := Position(2000)
	provider := NewFakeTime(position)
	var actual Position

	// Act
	c := make(chan Position, 1)
	provider.AfterFunc(distance, func() {
		c <- provider.Current()
	})
	provider.Update(newPosition)

	// Assert
	select {
	case actual = <-c:
		if actual < newPosition {
			t.Fatalf("The actual position was fired too early")
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("Timeout. the test exceed the expected duration")
	}
}

func TestFakeTimeAfterChan(t *testing.T) {
	// Arrange
	position := Position(1000)
	distance := Distance(100)
	newPosition := Position(2000)
	provider := NewFakeTime(position)
	var actual Position

	// Act
	c := make(chan Position, 1)
	provider.AfterChan(distance, c)
	provider.Update(newPosition)

	// Assert
	select {
	case actual = <-c:
		if actual < newPosition {
			t.Fatalf("The actual position was fired too early")
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("Timeout. the test exceed the expected duration")
	}
}

func TestFakeTimeAfterChanWithSleep(t *testing.T) {
	// Arrange
	position := Position(1000)
	distance := Distance(100)
	sleepDistance := Distance(200)
	provider := NewFakeTime(position)
	expected := Position(int64(position) + int64(distance))
	var actual Position

	// Act
	c := make(chan Position, 1)
	provider.AfterChan(distance, c)
	provider.Sleep(sleepDistance)

	// Assert
	select {
	case actual = <-c:
		if actual < expected {
			t.Fatalf("The actual position was fired too early")
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("Timeout. the test exceed the expected duration")
	}
}

func TestFakeTimeSince(t *testing.T) {
	// Arrange
	position := Position(1000)
	newPosition := Position(1200)
	provider := NewFakeTime(newPosition)
	expected := Distance(newPosition - position)
	var actual Distance

	// Act
	actual = provider.Since(position)

	// Assert
	if actual != expected {
		t.Fatalf("The actual distance mismatches the expected distance")
	}
}

func BenchmarkUpdate(b *testing.B) {
	// Arrange
	position := Position(0)
	provider := NewFakeTime(position)

	// Act
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			provider.Update(Position(j))
		}
	}
}
