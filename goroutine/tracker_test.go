package main

import (
	"context"
	"testing"
	"time"
)

func TestTracker(t *testing.T) {
	tr := NewTracker()
	go tr.Run()

	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")
	_ = tr.Event(context.Background(), "test")

	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr.Shutdown(ctx)
}
