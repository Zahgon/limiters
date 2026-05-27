package limiters

import (
	"context"
	"sync"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/redis/go-redis/v9"
)

// ConcurrentBufferBackend wraps the Add and Remove methods.
type ConcurrentBufferBackend interface {
	// Add adds the request with the given key to the buffer and returns the total number of requests in it.
	Add(ctx context.Context, key string) (int64, error)
	// Remove removes the request from the buffer.
	Remove(ctx context.Context, key string) error
}

// ConcurrentBuffer implements a limiter that allows concurrent requests up to the given capacity.
type ConcurrentBuffer struct {
	locker   DistLocker
	backend  ConcurrentBufferBackend
	logger   Logger
	capacity int64
	mu       sync.Mutex
}

// NewConcurrentBuffer creates a new ConcurrentBuffer instance.
func NewConcurrentBuffer(locker DistLocker, concurrentStateBackend ConcurrentBufferBackend, capacity int64, logger Logger) *ConcurrentBuffer {
	_ = "STUB: not implemented"
	return nil
}

// Limit puts the request identified by the key in a buffer.
func (c *ConcurrentBuffer) Limit(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Optimistically add the new request.

// Rollback the Add() operation.

// Done removes the request identified by the key from the buffer.
func (c *ConcurrentBuffer) Done(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// ConcurrentBufferInMemory is an in-memory implementation of ConcurrentBufferBackend.
type ConcurrentBufferInMemory struct {
	clock    Clock
	ttl      time.Duration
	mu       sync.Mutex
	registry *Registry
}

// NewConcurrentBufferInMemory creates a new instance of ConcurrentBufferInMemory.
// When the TTL of a key exceeds the key is removed from the buffer. This is needed in case if the process that added
// that key to the buffer did not call Done() for some reason.
func NewConcurrentBufferInMemory(registry *Registry, ttl time.Duration, clock Clock) *ConcurrentBufferInMemory {
	_ = "STUB: not implemented"
	return nil
}

// Add adds the request with the given key to the buffer and returns the total number of requests in it.
// It also removes the keys with expired TTL.
func (c *ConcurrentBufferInMemory) Add(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Remove removes the request from the buffer.
func (c *ConcurrentBufferInMemory) Remove(_ context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// ConcurrentBufferRedis implements ConcurrentBufferBackend in Redis.
type ConcurrentBufferRedis struct {
	clock Clock
	cli   redis.UniversalClient
	key   string
	ttl   time.Duration
}

// NewConcurrentBufferRedis creates a new instance of ConcurrentBufferRedis.
// When the TTL of a key exceeds the key is removed from the buffer. This is needed in case if the process that added
// that key to the buffer did not call Done() for some reason.
func NewConcurrentBufferRedis(cli redis.UniversalClient, key string, ttl time.Duration, clock Clock) *ConcurrentBufferRedis {
	_ = "STUB: not implemented"
	return nil
}

// Add adds the request with the given key to the sorted set in Redis and returns the total number of requests in it.
// It also removes the keys with expired TTL.
func (c *ConcurrentBufferRedis) Add(ctx context.Context, key string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Remove expired items.

// Remove removes the request identified by the key from the sorted set in Redis.
func (c *ConcurrentBufferRedis) Remove(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// ConcurrentBufferMemcached implements ConcurrentBufferBackend in Memcached.
type ConcurrentBufferMemcached struct {
	clock Clock
	cli   *memcache.Client
	key   string
	ttl   time.Duration
}

// NewConcurrentBufferMemcached creates a new instance of ConcurrentBufferMemcached.
// When the TTL of a key exceeds the key is removed from the buffer. This is needed in case if the process that added
// that key to the buffer did not call Done() for some reason.
func NewConcurrentBufferMemcached(cli *memcache.Client, key string, ttl time.Duration, clock Clock) *ConcurrentBufferMemcached {
	_ = "STUB: not implemented"
	return nil
}

type SortedSetNode struct {
	CreatedAt int64
	Value     string
}

// Add adds the request with the given key to the slice in Memcached and returns the total number of requests in it.
// It also removes the keys with expired TTL.
func (c *ConcurrentBufferMemcached) Add(ctx context.Context, element string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Remove removes the request identified by the key from the slice in Memcached.
func (c *ConcurrentBufferMemcached) Remove(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}
