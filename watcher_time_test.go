package axis

import (
    "testing"
    "time"
)


func TestTimeWatcherReset(t *testing.T) {
    // Arrange
    duration1 := time.Duration(2*time.Second)
    duration2 := time.Duration(4*time.Second)
    now := time.Now()
    timer := time.NewTimer(duration1)
    watcher := newTimeWatcher(timer)
    var actual bool
    var up time.Time

    // Act
    actual = watcher.Reset(Distance(duration2))

    // Assert
    select {
        case up = <-timer.C:
            if !actual {
                t.Fatalf("Cannot reset the time watcher")
            }
            if up.Before(now.Add(4*time.Second)) {
                t.Fatalf("The timer raise event before the expected date")
            }
        case <-time.After(5*time.Second):
            t.Fatalf("Timeout. the test exceed the expected duration")
    }
}

func TestTimeWatcherStop(t *testing.T) {
    // Arrange
    duration := time.Duration(2*time.Second)
    timer := time.NewTimer(duration)
    watcher := newTimeWatcher(timer)
    var actual bool

    // Act
    actual = watcher.Stop()

    // Assert
    if !actual {
        t.Fatalf("Cannot reset the time watcher")
    }
}
