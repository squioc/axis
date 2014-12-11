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
    date = time.Unix(int64(current), 0)

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
    position := <-c
    after := time.Unix(int64(position), 0)


    // Assert
    if expected.After(after) {
        t.Fatalf("The function After must fire the position after the minimal expected duration")
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
    position := <-c
    after := time.Unix(int64(position), 0)


    // Assert
    if watcher == nil {
        t.Fatalf("The function AfterChan must return a watcher")
    }
    if expected.After(after) {
        t.Fatalf("The function After must fire the position after the minimal expected duration")
    }
}
