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

// SlidingWindowIncrementer wraps the Increment method.
type SlidingWindowIncrementer interface {
	// Increment increments the request counter for the current window and returns the counter values for the previous
	// window and the current one.
	// TTL is the time duration before the next window.
	Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (prevCount, currCount int64, err error)
}

// SlidingWindow implements a Sliding Window rate limiting algorithm.
//
// It does not require a distributed lock and uses a minimum amount of memory, however it will disallow all the requests
// in case when a client is flooding the service with requests.
// It's the client's responsibility to handle the disallowed request and wait before making a new request again.
type SlidingWindow struct {
	backend  SlidingWindowIncrementer
	clock    Clock
	rate     time.Duration
	capacity int64
	epsilon  float64
}

// NewSlidingWindow creates a new instance of SlidingWindow.
// Capacity is the maximum amount of requests allowed per window.
// Rate is the window size.
// Epsilon is the max-allowed range of difference when comparing the current weighted number of requests with capacity.
func NewSlidingWindow(capacity int64, rate time.Duration, slidingWindowIncrementer SlidingWindowIncrementer, clock Clock, epsilon float64) *SlidingWindow {
	_ = "STUB: not implemented"
	return nil
}

// Limit returns the time duration to wait before the request can be processed.
// It returns ErrLimitExhausted if the request overflows the capacity.
func (s *SlidingWindow) Limit(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// "prev" and "curr" are capped at "s.capacity + s.epsilon" using math.Ceil to round up any fractional values,
// ensuring that in the worst case, "total" can be slightly greater than "s.capacity".

// If prev == 0.

// SlidingWindowInMemory is an in-memory implementation of SlidingWindowIncrementer.
type SlidingWindowInMemory struct {
	mu           sync.Mutex
	prevC, currC int64
	prevW, currW time.Time
}

// NewSlidingWindowInMemory creates a new instance of SlidingWindowInMemory.
func NewSlidingWindowInMemory() *SlidingWindowInMemory { _ = "STUB: not implemented"; return nil }

// Increment increments the current window's counter and returns the number of requests in the previous window and the
// current one.
func (s *SlidingWindowInMemory) Increment(ctx context.Context, prev, curr time.Time, _ time.Duration) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// SlidingWindowRedis implements SlidingWindow in Redis.
type SlidingWindowRedis struct {
	cli    redis.UniversalClient
	prefix string
}

// NewSlidingWindowRedis creates a new instance of SlidingWindowRedis.
func NewSlidingWindowRedis(cli redis.UniversalClient, prefix string) *SlidingWindowRedis {
	_ = "STUB: not implemented"
	return nil
}

// Increment increments the current window's counter in Redis and returns the number of requests in the previous window
// and the current one.
func (s *SlidingWindowRedis) Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// SlidingWindowMemcached implements SlidingWindow in Memcached.
type SlidingWindowMemcached struct {
	cli    *memcache.Client
	prefix string
}

// NewSlidingWindowMemcached creates a new instance of SlidingWindowMemcached.
func NewSlidingWindowMemcached(cli *memcache.Client, prefix string) *SlidingWindowMemcached {
	_ = "STUB: not implemented"
	return nil
}

// Increment increments the current window's counter in Memcached and returns the number of requests in the previous window
// and the current one.
func (s *SlidingWindowMemcached) Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// SlidingWindowDynamoDB implements SlidingWindow in DynamoDB.
type SlidingWindowDynamoDB struct {
	client       *dynamodb.Client
	partitionKey string
	tableProps   DynamoDBTableProperties
}

// NewSlidingWindowDynamoDB creates a new instance of SlidingWindowDynamoDB.
// PartitionKey is the key used to store all the this implementation in DynamoDB.
//
// TableProps describe the table that this backend should work with. This backend requires the following on the table:
// * SortKey
// * TTL.
func NewSlidingWindowDynamoDB(client *dynamodb.Client, partitionKey string, props DynamoDBTableProperties) *SlidingWindowDynamoDB {
	_ = "STUB: not implemented"
	return nil
}

// Increment increments the current window's counter in DynamoDB and returns the number of requests in the previous window
// and the current one.
func (s *SlidingWindowDynamoDB) Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// SlidingWindowCosmosDB implements SlidingWindow in Azure Cosmos DB.
type SlidingWindowCosmosDB struct {
	client       *azcosmos.ContainerClient
	partitionKey string
}

// NewSlidingWindowCosmosDB creates a new instance of SlidingWindowCosmosDB.
// PartitionKey is the key used to store all the this implementation in Cosmos.
func NewSlidingWindowCosmosDB(client *azcosmos.ContainerClient, partitionKey string) *SlidingWindowCosmosDB {
	_ = "STUB: not implemented"
	return nil
}

// Increment increments the current window's counter in Cosmos and returns the number of requests in the previous window
// and the current one.
func (s *SlidingWindowCosmosDB) Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// value exists and was updated

// Fallback to Read-Modify-Write
