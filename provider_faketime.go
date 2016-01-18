package axis

import (
	"sync"
)

// FakeTime simulate a time-based provider
type FakeTime struct {
	// position gives the position of the faketime provider
	position Position

	// stops represents specific points in the time
	// it calls the given function when it reaches them
	stops map[Position][]func()
	mu    sync.Mutex
}

// NewFakeTime creates a new faketime provider
func NewFakeTime(position Position) *FakeTime {
	return &FakeTime{
		position: position,
		stops:    make(map[Position][]func()),
	}
}

// Current gets the current position of the provider
func (f *FakeTime) Current() Position {
	return f.position
}

// Sleep pauses the provider for the given distance
func (f *FakeTime) Sleep(distance Distance) {
	f.Update(addDistance(f.position, distance))
}

// After simulates a wait for the given distance to elapse
// and then sends the new position on the returned channel
func (f *FakeTime) After(distance Distance) <-chan Position {
	c := make(chan Position, 1)
	f.AfterChan(distance, c)
	return c
}

// AfterFunc simulates a wait for the given distance to elapse
// and then calls the callback
func (f *FakeTime) AfterFunc(distance Distance, callback func()) Watcher {
	f.mu.Lock()
	until := addDistance(f.position, distance)
	f.stops[until] = append(f.stops[until], callback)
	f.mu.Unlock()
	return &FakeTimeWatcher{canReset: true, canStop: true}
}

// AfterChan simulates a wait for the given distance to elapse
// and then sends the new position on the given channel
func (f *FakeTime) AfterChan(distance Distance, channel chan Position) Watcher {
	return f.AfterFunc(distance, func() {
		channel <- f.Current()
	})
}

// Since returns the distance traveled since position
func (f *FakeTime) Since(position Position) Distance {
	return Distance(f.Current() - position)
}

// Update sets the current position of the provider
func (f *FakeTime) Update(position Position) {
	f.position = position

	f.mu.Lock()
	for k, v := range f.stops {
		if k < f.position {
			for _, c := range v {
				c()
			}
			delete(f.stops, k)
		}
	}
	f.mu.Unlock()
}
