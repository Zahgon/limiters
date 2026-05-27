package limiters

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/redis/go-redis/v9"
)

// FixedWindowIncrementer wraps the Increment method.
type FixedWindowIncrementer interface {
	// Increment increments the request counter for the window and returns the counter value.
	// TTL is the time duration before the next window.
	Increment(ctx context.Context, window time.Time, ttl time.Duration) (int64, error)
}

// FixedWindow implements a Fixed Window rate limiting algorithm.
//
// Simple and memory efficient algorithm that does not need a distributed lock.
// However it may be lenient when there are many requests around the boundary between 2 adjacent windows.
type FixedWindow struct {
	backend  FixedWindowIncrementer
	clock    Clock
	rate     time.Duration
	capacity int64
	mu       sync.Mutex
	window   time.Time
	overflow bool
}

// NewFixedWindow creates a new instance of FixedWindow.
// Capacity is the maximum amount of requests allowed per window.
// Rate is the window size.
func NewFixedWindow(capacity int64, rate time.Duration, fixedWindowIncrementer FixedWindowIncrementer, clock Clock) *FixedWindow {
	_ = "STUB: not implemented"
	return nil
}

// Limit returns the time duration to wait before the request can be processed.
// It returns ErrLimitExhausted if the request overflows the window's capacity.
func (f *FixedWindow) Limit(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// If the window is already overflowed don't increment the counter.

// FixedWindowInMemory is an in-memory implementation of FixedWindowIncrementer.
type FixedWindowInMemory struct {
	mu     sync.Mutex
	c      int64
	window time.Time
}

// NewFixedWindowInMemory creates a new instance of FixedWindowInMemory.
func NewFixedWindowInMemory() *FixedWindowInMemory { _ = "STUB: not implemented"; return nil }

// Increment increments the window's counter.
func (f *FixedWindowInMemory) Increment(ctx context.Context, window time.Time, _ time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FixedWindowRedis implements FixedWindow in Redis.
type FixedWindowRedis struct {
	cli    redis.UniversalClient
	prefix string
}

// NewFixedWindowRedis returns a new instance of FixedWindowRedis.
// Prefix is the key prefix used to store all the keys used in this implementation in Redis.
func NewFixedWindowRedis(cli redis.UniversalClient, prefix string) *FixedWindowRedis {
	_ = "STUB: not implemented"
	return nil
}

// Increment increments the window's counter in Redis.
func (f *FixedWindowRedis) Increment(ctx context.Context, window time.Time, ttl time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FixedWindowMemcached implements FixedWindow in Memcached.
type FixedWindowMemcached struct {
	cli    *memcache.Client
	prefix string
}

// NewFixedWindowMemcached returns a new instance of FixedWindowMemcached.
// Prefix is the key prefix used to store all the keys used in this implementation in Memcached.
func NewFixedWindowMemcached(cli *memcache.Client, prefix string) *FixedWindowMemcached {
	_ = "STUB: not implemented"
	return nil
}

// Increment increments the window's counter in Memcached.
func (f *FixedWindowMemcached) Increment(ctx context.Context, window time.Time, ttl time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FixedWindowDynamoDB implements FixedWindow in DynamoDB.
type FixedWindowDynamoDB struct {
	client       *dynamodb.Client
	partitionKey string
	tableProps   DynamoDBTableProperties
}

// NewFixedWindowDynamoDB creates a new instance of FixedWindowDynamoDB.
// PartitionKey is the key used to store all the this implementation in DynamoDB.
//
// TableProps describe the table that this backend should work with. This backend requires the following on the table:
// * SortKey
// * TTL.
func NewFixedWindowDynamoDB(client *dynamodb.Client, partitionKey string, props DynamoDBTableProperties) *FixedWindowDynamoDB {
	_ = "STUB: not implemented"
	return nil
}

type contextKey int

var fixedWindowDynamoDBPartitionKey contextKey

// NewFixedWindowDynamoDBContext creates a context for FixedWindowDynamoDB with a partition key.
//
// This context can be used to control the partition key per-request.
//
// Deprecated: NewFixedWindowDynamoDBContext is deprecated and will be removed in future versions.
// Separate FixedWindow rate limiters should be used for different partition keys instead.
// Consider using the `Registry` to manage multiple FixedWindow instances with different partition keys.
func NewFixedWindowDynamoDBContext(ctx context.Context, partitionKey string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

const (
	fixedWindowDynamoDBUpdateExpression = "SET #C = if_not_exists(#C, :def) + :inc, #TTL = :ttl"
	dynamodbWindowCountKey              = "Count"
)

// Increment increments the window's counter in DynamoDB.
func (f *FixedWindowDynamoDB) Increment(ctx context.Context, window time.Time, ttl time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FixedWindowCosmosDB implements FixedWindow in CosmosDB.
type FixedWindowCosmosDB struct {
	client       *azcosmos.ContainerClient
	partitionKey string
}

// NewFixedWindowCosmosDB creates a new instance of FixedWindowCosmosDB.
// PartitionKey is the key used for partitioning data into multiple partitions.
func NewFixedWindowCosmosDB(client *azcosmos.ContainerClient, partitionKey string) *FixedWindowCosmosDB {
	_ = "STUB: not implemented"
	return nil
}

func (f *FixedWindowCosmosDB) Increment(ctx context.Context, window time.Time, ttl time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// value exists and was updated

// Try to CreateItem. If it fails with Conflict, it means the item exists (and Patch failed with 400), so we fallback to RMW.

// Fallback to Read-Modify-Write
