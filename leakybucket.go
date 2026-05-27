package limiters

import (
	"context"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// LeakyBucketState represents the state of a LeakyBucket.
type LeakyBucketState struct {
	// Last is the Unix timestamp in nanoseconds of the most recent request.
	Last int64
}

// IzZero returns true if the bucket state is zero valued.
func (s LeakyBucketState) IzZero() bool {
	_ = "STUB: not implemented"

	// LeakyBucketStateBackend interface encapsulates the logic of retrieving and persisting the state of a LeakyBucket.
	return false
}

type LeakyBucketStateBackend interface {
	// State gets the current state of the LeakyBucket.
	State(ctx context.Context) (LeakyBucketState, error)
	// SetState sets (persists) the current state of the LeakyBucket.
	SetState(ctx context.Context, state LeakyBucketState) error
	// Reset resets (persists) the current state of the LeakyBucket.
	Reset(ctx context.Context) error
}

// LeakyBucket implements the https://en.wikipedia.org/wiki/Leaky_bucket#As_a_queue algorithm.
type LeakyBucket struct {
	locker  DistLocker
	backend LeakyBucketStateBackend
	clock   Clock
	logger  Logger
	// Capacity is the maximum allowed number of tokens in the bucket.
	capacity int64
	// Rate is the output rate: 1 request per the rate duration (in nanoseconds).
	rate int64
	mu   sync.Mutex
}

// NewLeakyBucket creates a new instance of LeakyBucket.
func NewLeakyBucket(capacity int64, rate time.Duration, locker DistLocker, leakyBucketStateBackend LeakyBucketStateBackend, clock Clock, logger Logger) *LeakyBucket {
	_ = "STUB: not implemented"
	return nil
}

// Limit returns the time duration to wait before the request can be processed.
// It returns ErrLimitExhausted if the request overflows the bucket's capacity. In this case the returned duration
// means how long it would have taken to wait for the request to be processed if the bucket was not overflowed.
func (t *LeakyBucket) Limit(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// The queue has requests in it: move the current request to the last position + 1.

// The queue is empty.
// The offset is the duration to wait in case the last request happened less than rate duration ago.

// Reset resets the bucket.
func (t *LeakyBucket) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// LeakyBucketInMemory is an in-memory implementation of LeakyBucketStateBackend.
type LeakyBucketInMemory struct {
	state LeakyBucketState
}

// NewLeakyBucketInMemory creates a new instance of LeakyBucketInMemory.
func NewLeakyBucketInMemory() *LeakyBucketInMemory { _ = "STUB: not implemented"; return nil }

// State gets the current state of the bucket.
func (l *LeakyBucketInMemory) State(ctx context.Context) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	return *new(LeakyBucketState), nil
}

// SetState sets the current state of the bucket.
func (l *LeakyBucketInMemory) SetState(ctx context.Context, state LeakyBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset resets the current state of the bucket.
func (l *LeakyBucketInMemory) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	etcdKeyLBLease = "lease"
	etcdKeyLBLast  = "last"
)

// LeakyBucketEtcd is an etcd implementation of a LeakyBucketStateBackend.
// See the TokenBucketEtcd description for the details on etcd usage.
type LeakyBucketEtcd struct {
	// prefix is the etcd key prefix.
	prefix      string
	cli         *clientv3.Client
	leaseID     clientv3.LeaseID
	ttl         time.Duration
	raceCheck   bool
	lastVersion int64
}

// NewLeakyBucketEtcd creates a new LeakyBucketEtcd instance.
// Prefix is used as an etcd key prefix for all keys stored in etcd by this algorithm.
// TTL is a TTL of the etcd lease used to store all the keys.
//
// If raceCheck is true and the keys in etcd are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewLeakyBucketEtcd(cli *clientv3.Client, prefix string, ttl time.Duration, raceCheck bool) *LeakyBucketEtcd {
	_ = "STUB: not implemented"
	return nil
}

// State gets the bucket's current state from etcd.
// If there is no state available in etcd then the initial bucket's state is returned.
func (l *LeakyBucketEtcd) State(ctx context.Context) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	// Reset the lease ID as it will be updated by the successful Get operation below.
	return *new(LeakyBucketState), nil
}

// Get all the keys under the prefix in a single request.

// Ignore lease when there is no expiration

// createLease creates a new lease in etcd and updates the t.leaseID value.
func (l *LeakyBucketEtcd) createLease(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// save saves the state to etcd using the existing lease.
func (l *LeakyBucketEtcd) save(ctx context.Context, state LeakyBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Put the keys only if they have not been modified since the most recent read.

// SetState updates the state of the bucket in etcd.
func (l *LeakyBucketEtcd) SetState(ctx context.Context, state LeakyBucketState) error {
	_ = "STUB: not implemented"

	// Avoid maintaining the lease when it has no TTL
	return nil
}

// Lease does not exist, create one.

// No need to send KeepAlive for the newly creates lease: save the state immediately.

// Send the KeepAlive request to extend the existing lease.

// Create a new lease since the current one has expired.

// Reset resets the state of the bucket in etcd.
func (l *LeakyBucketEtcd) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Deprecated: These legacy keys will be removed in a future version.
// The state is now stored in a single JSON document under the "state" key.
const (
	redisKeyLBLast    = "last"
	redisKeyLBVersion = "version"
)

// LeakyBucketRedis is a Redis implementation of a LeakyBucketStateBackend.
type LeakyBucketRedis struct {
	cli         redis.UniversalClient
	prefix      string
	ttl         time.Duration
	raceCheck   bool
	lastVersion int64
}

// NewLeakyBucketRedis creates a new LeakyBucketRedis instance.
// Prefix is the key prefix used to store all the keys used in this implementation in Redis.
// TTL is the TTL of the stored keys.
//
// If raceCheck is true and the keys in Redis are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewLeakyBucketRedis(cli redis.UniversalClient, prefix string, ttl time.Duration, raceCheck bool) *LeakyBucketRedis {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: Legacy format support will be removed in a future version.
func (t *LeakyBucketRedis) oldState(ctx context.Context) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	return *new(LeakyBucketState), nil
}

// Keys don't exist, return an empty state.

// State gets the bucket's state from Redis.
func (t *LeakyBucketRedis) State(ctx context.Context) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	return *new(LeakyBucketState), nil
}

// reset in a case of returning an empty LeakyBucketState

// Try new format

// SetState updates the state in Redis.
func (t *LeakyBucketRedis) SetState(ctx context.Context, state LeakyBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset resets the state in Redis.
func (t *LeakyBucketRedis) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// LeakyBucketMemcached is a Memcached implementation of a LeakyBucketStateBackend.
type LeakyBucketMemcached struct {
	cli       *memcache.Client
	key       string
	ttl       time.Duration
	raceCheck bool
	casId     uint64
}

// NewLeakyBucketMemcached creates a new LeakyBucketMemcached instance.
// Key is the key used to store all the keys used in this implementation in Memcached.
// TTL is the TTL of the stored keys.
//
// If raceCheck is true and the keys in Memcached are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewLeakyBucketMemcached(cli *memcache.Client, key string, ttl time.Duration, raceCheck bool) *LeakyBucketMemcached {
	_ = "STUB: not implemented"
	return nil
}

// State gets the bucket's state from Memcached.
func (t *LeakyBucketMemcached) State(ctx context.Context) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	return *new(LeakyBucketState), nil
}

// Keys don't exist, return an empty state.

// SetState updates the state in Memcached.
// The provided fencing token is checked on the Memcached side before saving the keys.
func (t *LeakyBucketMemcached) SetState(ctx context.Context, state LeakyBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// If the value is over 30 days, it treats it as UNIX timestamp.

// Memcached supports expiration in seconds. It's more precise way.

// Reset resets the state in Memcached.
func (t *LeakyBucketMemcached) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// LeakyBucketDynamoDB is a DyanamoDB implementation of a LeakyBucketStateBackend.
type LeakyBucketDynamoDB struct {
	client        *dynamodb.Client
	tableProps    DynamoDBTableProperties
	partitionKey  string
	ttl           time.Duration
	raceCheck     bool
	latestVersion int64
	keys          map[string]types.AttributeValue
}

// NewLeakyBucketDynamoDB creates a new LeakyBucketDynamoDB instance.
// PartitionKey is the key used to store all the this implementation in DynamoDB.
//
// TableProps describe the table that this backend should work with. This backend requires the following on the table:
// * TTL
//
// TTL is the TTL of the stored item.
//
// If raceCheck is true and the item in DynamoDB are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewLeakyBucketDynamoDB(client *dynamodb.Client, partitionKey string, tableProps DynamoDBTableProperties, ttl time.Duration, raceCheck bool) *LeakyBucketDynamoDB {
	_ = "STUB: not implemented"
	return nil
}

// State gets the bucket's state from DynamoDB.
func (t *LeakyBucketDynamoDB) State(ctx context.Context) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	return *new(LeakyBucketState), nil
}

// SetState updates the state in DynamoDB.
func (t *LeakyBucketDynamoDB) SetState(ctx context.Context, state LeakyBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset resets the state in DynamoDB.
func (t *LeakyBucketDynamoDB) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	dynamodbBucketRaceConditionExpression = "Version <= :version"
	dynamoDBBucketLastKey                 = "Last"
	dynamoDBBucketVersionKey              = "Version"
)

func (t *LeakyBucketDynamoDB) getPutItemInputFromState(state LeakyBucketState) *dynamodb.PutItemInput {
	_ = "STUB: not implemented"
	return nil
}

func (t *LeakyBucketDynamoDB) getGetItemInput() *dynamodb.GetItemInput {
	_ = "STUB: not implemented"
	return nil
}

func (t *LeakyBucketDynamoDB) loadStateFromDynamoDB(resp *dynamodb.GetItemOutput) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	return *new(LeakyBucketState), nil
}

// CosmosDBLeakyBucketItem represents a document in CosmosDB for LeakyBucket.
type CosmosDBLeakyBucketItem struct {
	ID           string           `json:"id"`
	PartitionKey string           `json:"partitionKey"`
	State        LeakyBucketState `json:"state"`
	Version      int64            `json:"version"`
	TTL          int64            `json:"ttl,omitempty"`
}

// LeakyBucketCosmosDB is a CosmosDB implementation of a LeakyBucketStateBackend.
type LeakyBucketCosmosDB struct {
	client        *azcosmos.ContainerClient
	partitionKey  string
	id            string
	ttl           time.Duration
	raceCheck     bool
	latestVersion int64
}

// NewLeakyBucketCosmosDB creates a new LeakyBucketCosmosDB instance.
// PartitionKey is the key used to store all the implementation in CosmosDB.
// TTL is the TTL of the stored item.
//
// If raceCheck is true and the item in CosmosDB is modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewLeakyBucketCosmosDB(client *azcosmos.ContainerClient, partitionKey string, ttl time.Duration, raceCheck bool) *LeakyBucketCosmosDB {
	_ = "STUB: not implemented"
	return nil
}

func (t *LeakyBucketCosmosDB) State(ctx context.Context) (LeakyBucketState, error) {
	_ = "STUB: not implemented"
	return *new(LeakyBucketState), nil
}

func (t *LeakyBucketCosmosDB) SetState(ctx context.Context, state LeakyBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *LeakyBucketCosmosDB) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
