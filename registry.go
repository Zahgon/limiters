package limiters

import (
	"sync"
	"time"
)

// pqItem is an item in the priority queue.
type pqItem struct {
	value any
	exp   time.Time
	index int
	key   string
}

// gcPq is a priority queue.
type gcPq []*pqItem

func (pq gcPq) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq gcPq) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq gcPq) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *gcPq) Push(x any) { _ = "STUB: not implemented"; return }

func (pq *gcPq) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// for safety

// Registry is a thread-safe garbage-collectable registry of values.
type Registry struct {
	// Guards all the fields below it.
	mx sync.Mutex
	pq *gcPq
	m  map[string]*pqItem
}

// NewRegistry creates a new instance of Registry.
func NewRegistry() *Registry { _ = "STUB: not implemented"; return nil }

// GetOrCreate gets an existing value by key and updates its expiration time.
// If the key lookup fails it creates a new value by calling the provided value closure and puts it on the queue.
func (r *Registry) GetOrCreate(key string, value func() any, ttl time.Duration, now time.Time) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Update the expiration time.

// DeleteExpired deletes expired items from the registry and returns the number of deleted items.
func (r *Registry) DeleteExpired(now time.Time) int { _ = "STUB: not implemented"; return 0 }

// Delete deletes an item from the registry.
func (r *Registry) Delete(key string) { _ = "STUB: not implemented"; return }

// Exists returns true if an item with the given key exists in the registry.
func (r *Registry) Exists(key string) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of items in the registry.
func (r *Registry) Len() int { _ = "STUB: not implemented"; return 0 }
