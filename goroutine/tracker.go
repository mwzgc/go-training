package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Go")
}

// NewTracker ...
func NewTracker() *Tracker {
	return &Tracker{
		ch:   make(chan string, 10),
		stop: make(chan struct{}),
	}
}

// Tracker ...
type Tracker struct {
	ch   chan string
	stop chan struct{}
}

// Event ...
func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Run ...
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

// Shutdown ...
func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)

	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}
