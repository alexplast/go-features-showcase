package features

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Function that demonstrates goroutines and channels.
func LongRunningTask(c chan string) {
	log.Info("Long running task started...")
	time.Sleep(2 * time.Second)
	c <- "Long running task finished!"
}

// Function to demonstrate the select statement.
func DemonstrateSelect() {
	log.Info("\n--- Select ---")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			log.Info("received", msg1)
		case msg2 := <-c2:
			log.Info("received", msg2)
		}
	}
}

// A SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func DemonstrateMutex() {
	log.Info("\n--- Mutex ---")
	c := SafeCounter{v: make(map[string]int)}
	for range 1000 {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	log.Info("Counter value:", c.Value("somekey"))
}

func DemonstrateContext(ctx context.Context) {
	log.Info("\n--- Context ---")
	select {
	case <-time.After(2 * time.Second):
		log.Info("overslept")
	case <-ctx.Done():
		log.Info("Context cancelled:", ctx.Err())
	}
}
