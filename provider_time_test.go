package axis

import (
	"testing"
	"time"
)

func TestTimeCurrent(t *testing.T) {
	// Arrange
	provider := &Time{}
	var current Position
	var date time.Time

	// Act
	now := time.Unix(time.Now().Unix(), 0)
	current = provider.Current()
	date = time.Unix(int64(current) / 1000, int64(current) % 1000)

	// Assert
	if now.After(date) {
		t.Fatalf("The current position must refer to the present date")
	}
}

func TestTimeSleep(t *testing.T) {
	// Arrange
	provider := &Time{}
	duration := time.Second

	// Act
	before := time.Now()
	expected := before.Add(duration)
	provider.Sleep(Distance(duration))
	after := time.Now()

	// Assert
	if expected.After(after) {
		t.Fatalf("The duration of the call of the function Sleep mismatch the minimal expected duration")
	}
}

func TestTimeAfter(t *testing.T) {
	// Arrange
	provider := &Time{}
	duration := time.Second

	// Act
	before := time.Unix(time.Now().Unix(), 0)
	expected := before.Add(duration)
	c := provider.After(Distance(duration))

	// Assert
	select {
	case position := <-c:
		after := time.Unix(int64(position), 0)
		if expected.After(after) {
			t.Fatalf("The function After must fire the position after the minimal expected duration")
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("Timeout. the test exceed the expected duration")
	}
}

func TestTimeAfterFunc(t *testing.T) {
	// Arrange
	provider := &Time{}
	duration := time.Second
	c := make(chan Position, 1)

	// Act
	before := time.Unix(time.Now().Unix(), 0)
	expected := before.Add(duration)
	watcher := provider.AfterFunc(Distance(duration), func() {
		c <- provider.Current()
	})

	// Assert
	select {
	case position := <-c:
		after := time.Unix(int64(position), 0)
		if watcher == nil {
			t.Fatalf("The function AfterChan must return a watcher")
		}
		if expected.After(after) {
			t.Fatalf("The function After must fire the position after the minimal expected duration")
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("Timeout. the test exceed the expected duration")
	}
}

func TestTimeAfterChan(t *testing.T) {
	// Arrange
	provider := &Time{}
	duration := time.Second
	c := make(chan Position, 1)

	// Act
	before := time.Unix(time.Now().Unix(), 0)
	expected := before.Add(duration)
	watcher := provider.AfterChan(Distance(duration), c)

	// Assert
	select {
	case position := <-c:
		after := time.Unix(int64(position), 0)
		if watcher == nil {
			t.Fatalf("The function AfterChan must return a watcher")
		}
		if expected.After(after) {
			t.Fatalf("The function After must fire the position after the minimal expected duration")
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("Timeout. the test exceed the expected duration")
	}
}

func TestTimeSince(t *testing.T) {
	// Arrange
	provider := &Time{}
	expected := Distance(200)
	position := Position(time.Now().UnixNano() - int64(expected))
	var actual Distance

	// Act
	actual = provider.Since(position)

	// Assert
	if actual < expected {
		t.Fatalf("The actual distance is lesser than the expected distance")
	}
}
