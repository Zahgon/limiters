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
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// TokenBucketState represents a state of a token bucket.
type TokenBucketState struct {
	// Last is the last time the state was updated (Unix timestamp in nanoseconds).
	Last int64
	// Available is the number of available tokens in the bucket.
	Available int64
}

// isZero returns true if the bucket state is zero valued.
func (s TokenBucketState) isZero() bool { _ = "STUB: not implemented"; return false }

// TokenBucketStateBackend interface encapsulates the logic of retrieving and persisting the state of a TokenBucket.
type TokenBucketStateBackend interface {
	// State gets the current state of the TokenBucket.
	State(ctx context.Context) (TokenBucketState, error)
	// SetState sets (persists) the current state of the TokenBucket.
	SetState(ctx context.Context, state TokenBucketState) error
	// Reset resets (persists) the current state of the TokenBucket.
	Reset(ctx context.Context) error
}

// TokenBucket implements the https://en.wikipedia.org/wiki/Token_bucket algorithm.
type TokenBucket struct {
	locker  DistLocker
	backend TokenBucketStateBackend
	clock   Clock
	logger  Logger
	// refillRate is the tokens refill rate (1 token per duration).
	refillRate time.Duration
	// capacity is the bucket's capacity.
	capacity int64
	mu       sync.Mutex
}

// NewTokenBucket creates a new instance of TokenBucket.
func NewTokenBucket(capacity int64, refillRate time.Duration, locker DistLocker, tokenBucketStateBackend TokenBucketStateBackend, clock Clock, logger Logger) *TokenBucket {
	_ = "STUB: not implemented"
	return nil
}

// takeMinMax takes between minTokens and maxTokens tokens from the bucket, depending on availability.
//
// It returns the number of tokens actually taken, zero duration, and nil error if at least minTokens are available.
//
// If fewer than minTokens are available, it returns 0 taken, duration indicating how long to wait before retrying, and ErrLimitExhausted error.
// The wait duration is computed based on the refill rate and the deficit to reach minTokens.
func (t *TokenBucket) takeMinMax(ctx context.Context, minTokens, maxTokens int64) (int64, time.Duration, error) {
	_ = "STUB: not implemented"
	return 0, *new(time.Duration), nil
}

// Initially the bucket is full.

// Refill the bucket.

// Take as many tokens between minTokens and maxTokens as possible.

// Take the tokens from the bucket.

// TakeMax takes up to maxTokens tokens from the bucket, depending on availability.
//
// It returns the number of tokens actually taken, and error in case internal action will fail.
func (t *TokenBucket) TakeMax(ctx context.Context, tokens int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Take takes tokens from the bucket.
//
// It returns a zero duration and a nil error if the bucket has sufficient amount of tokens.
//
// It returns ErrLimitExhausted if the amount of available tokens is less than requested. In this case the returned
// duration is the amount of time to wait to retry the request.
func (t *TokenBucket) Take(ctx context.Context, tokens int64) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Limit takes 1 token from the bucket.
func (t *TokenBucket) Limit(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *

	// Reset resets the bucket.
	new(time.Duration), nil
}

func (t *TokenBucket) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TokenBucketInMemory is an in-memory implementation of TokenBucketStateBackend.
//
// The state is not shared nor persisted so it won't survive restarts or failures.
// Due to the local nature of the state the rate at which some endpoints are accessed can't be reliably predicted or
// limited.
//
// Although it can be used as a global rate limiter with a round-robin load-balancer.
type TokenBucketInMemory struct {
	state TokenBucketState
}

// NewTokenBucketInMemory creates a new instance of TokenBucketInMemory.
func NewTokenBucketInMemory() *TokenBucketInMemory { _ = "STUB: not implemented"; return nil }

// State returns the current bucket's state.
func (t *TokenBucketInMemory) State(ctx context.Context) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	return *new(TokenBucketState), nil
}

// SetState sets the current bucket's state.
func (t *TokenBucketInMemory) SetState(ctx context.Context, state TokenBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset resets the current bucket's state.
func (t *TokenBucketInMemory) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	etcdKeyTBLease     = "lease"
	etcdKeyTBAvailable = "available"
	etcdKeyTBLast      = "last"
)

// TokenBucketEtcd is an etcd implementation of a TokenBucketStateBackend.
//
// See https://github.com/etcd-io/etcd/blob/master/Documentation/learning/data_model.md
//
// etcd is designed to reliably store infrequently updated data, thus it should only be used for the API endpoints which
// are accessed less frequently than it can be processed by the rate limiter.
//
// Aggressive compaction and defragmentation has to be enabled in etcd to prevent the size of the storage
// to grow indefinitely: every change of the state of the bucket (every access) will create a new revision in etcd.
//
// It probably makes it impractical for the high load cases, but can be used to reliably and precisely rate limit an
// access to the business critical endpoints where each access must be reliably logged.
type TokenBucketEtcd struct {
	// prefix is the etcd key prefix.
	prefix      string
	cli         *clientv3.Client
	leaseID     clientv3.LeaseID
	ttl         time.Duration
	raceCheck   bool
	lastVersion int64
}

// NewTokenBucketEtcd creates a new TokenBucketEtcd instance.
// Prefix is used as an etcd key prefix for all keys stored in etcd by this algorithm.
// TTL is a TTL of the etcd lease in seconds used to store all the keys: all the keys are automatically deleted after
// the TTL expires.
//
// If raceCheck is true and the keys in etcd are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
// It does not add any significant overhead as it can be trivially checked on etcd side before updating the keys.
func NewTokenBucketEtcd(cli *clientv3.Client, prefix string, ttl time.Duration, raceCheck bool) *TokenBucketEtcd {
	_ = "STUB: not implemented"
	return nil
}

// etcdKey returns a full etcd key from the provided key and prefix.
func etcdKey(prefix, key string) string { _ = "STUB: not implemented"; return "" }

// parseEtcdInt64 parses the etcd value into int64.
func parseEtcdInt64(kv *mvccpb.KeyValue) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func incPrefix(p string) string { _ = "STUB: not implemented"; return "" }

// State gets the bucket's current state from etcd.
// If there is no state available in etcd then the initial bucket's state is returned.
func (t *TokenBucketEtcd) State(ctx context.Context) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	// Get all the keys under the prefix in a single request.
	return *new(TokenBucketState), nil
}

// State not found, return zero valued state.

// Ignore lease when there is no expiration

// createLease creates a new lease in etcd and updates the t.leaseID value.
func (t *TokenBucketEtcd) createLease(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// save saves the state to etcd using the existing lease and the fencing token.
func (t *TokenBucketEtcd) save(ctx context.Context, state TokenBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Put the keys only if they have not been modified since the most recent read.

// SetState updates the state of the bucket.
func (t *TokenBucketEtcd) SetState(ctx context.Context, state TokenBucketState) error {
	_ = "STUB: not implemented"

	// Avoid maintaining the lease when it has no TTL
	return nil
}

// Lease does not exist, create one.

// No need to send KeepAlive for the newly created lease: save the state immediately.

// Send the KeepAlive request to extend the existing lease.

// Create a new lease since the current one has expired.

// Reset resets the state of the bucket.
func (t *TokenBucketEtcd) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Deprecated: These legacy keys will be removed in a future version.
// The state is now stored in a single JSON document under the "state" key.
const (
	redisKeyTBAvailable = "available"
	redisKeyTBLast      = "last"
	redisKeyTBVersion   = "version"
)

// If we do use cluster client and if the cluster is large enough, it is possible that when accessing multiple keys
// in leaky bucket or token bucket, these keys might go different slots and it will fail with error message
// `CROSSSLOT Keys in request don't hash to the same slot`. Adding hash tags in redisKey will force them into the
// same slot for keys with the same prefix.
//
// https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/#hash-tags
func redisKey(prefix, key string) string { _ = "STUB: not implemented"; return "" }

// TokenBucketRedis is a Redis implementation of a TokenBucketStateBackend.
//
// Redis is an in-memory key-value data storage which also supports persistence.
// It is a better choice for high load cases than etcd as it does not keep old values of the keys thus does not need
// the compaction/defragmentation.
//
// Although depending on a persistence and a cluster configuration some data might be lost in case of a failure
// resulting in an under-limiting the accesses to the service.
type TokenBucketRedis struct {
	cli         redis.UniversalClient
	prefix      string
	ttl         time.Duration
	raceCheck   bool
	lastVersion int64
}

// NewTokenBucketRedis creates a new TokenBucketRedis instance.
// Prefix is the key prefix used to store all the keys used in this implementation in Redis.
// TTL is the TTL of the stored keys.
//
// If raceCheck is true and the keys in Redis are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
// This adds an extra overhead since a Lua script has to be executed on the Redis side which locks the entire database.
func NewTokenBucketRedis(cli redis.UniversalClient, prefix string, ttl time.Duration, raceCheck bool) *TokenBucketRedis {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: Legacy format support will be removed in a future version.
func (t *TokenBucketRedis) oldState(ctx context.Context) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	return *new(TokenBucketState), nil
}

// reset in a case of returning an empty TokenBucketState

// Keys don't exist, return the initial state.

// State gets the bucket's state from Redis.
func (t *TokenBucketRedis) State(ctx context.Context) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	return *new(TokenBucketState), nil
}

// reset in a case of returning an empty TokenBucketState

// Try new format

// SetState updates the state in Redis.
func (t *TokenBucketRedis) SetState(ctx context.Context, state TokenBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset resets the state in Redis.
func (t *TokenBucketRedis) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TokenBucketMemcached is a Memcached implementation of a TokenBucketStateBackend.
//
// Memcached is a distributed memory object caching system.
type TokenBucketMemcached struct {
	cli       *memcache.Client
	key       string
	ttl       time.Duration
	raceCheck bool
	casId     uint64
}

// NewTokenBucketMemcached creates a new TokenBucketMemcached instance.
// Key is the key used to store all the keys used in this implementation in Memcached.
// TTL is the TTL of the stored keys.
//
// If raceCheck is true and the keys in Memcached are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
// This adds an extra overhead since a Lua script has to be executed on the Memcached side which locks the entire database.
func NewTokenBucketMemcached(cli *memcache.Client, key string, ttl time.Duration, raceCheck bool) *TokenBucketMemcached {
	_ = "STUB: not implemented"
	return nil
}

// State gets the bucket's state from Memcached.
func (t *TokenBucketMemcached) State(ctx context.Context) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	return *new(TokenBucketState), nil
}

// Keys don't exist, return the initial state.

// SetState updates the state in Memcached.
func (t *TokenBucketMemcached) SetState(ctx context.Context, state TokenBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// If the value is over 30 days, it treats it as UNIX timestamp.

// Memcached supports expiration in seconds. It's more precise way.

// Reset resets the state in Memcached.
func (t *TokenBucketMemcached) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Override casId to 0 to Set instead of CompareAndSwap in SetState

// TokenBucketDynamoDB is a DynamoDB implementation of a TokenBucketStateBackend.
type TokenBucketDynamoDB struct {
	client        *dynamodb.Client
	tableProps    DynamoDBTableProperties
	partitionKey  string
	ttl           time.Duration
	raceCheck     bool
	latestVersion int64
	keys          map[string]types.AttributeValue
}

// NewTokenBucketDynamoDB creates a new TokenBucketDynamoDB instance.
// PartitionKey is the key used to store all the this implementation in DynamoDB.
//
// TableProps describe the table that this backend should work with. This backend requires the following on the table:
// * TTL
//
// TTL is the TTL of the stored item.
//
// If raceCheck is true and the item in DynamoDB are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewTokenBucketDynamoDB(client *dynamodb.Client, partitionKey string, tableProps DynamoDBTableProperties, ttl time.Duration, raceCheck bool) *TokenBucketDynamoDB {
	_ = "STUB: not implemented"
	return nil
}

// State gets the bucket's state from DynamoDB.
func (t *TokenBucketDynamoDB) State(ctx context.Context) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	return *new(TokenBucketState), nil
}

// SetState updates the state in DynamoDB.
func (t *TokenBucketDynamoDB) SetState(ctx context.Context, state TokenBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset resets the state in DynamoDB.
func (t *TokenBucketDynamoDB) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

const dynamoDBBucketAvailableKey = "Available"

func (t *TokenBucketDynamoDB) getGetItemInput() *dynamodb.GetItemInput {
	_ = "STUB: not implemented"
	return nil
}

func (t *TokenBucketDynamoDB) getPutItemInputFromState(state TokenBucketState) *dynamodb.PutItemInput {
	_ = "STUB: not implemented"
	return nil
}

func (t *TokenBucketDynamoDB) loadStateFromDynamoDB(resp *dynamodb.GetItemOutput) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	return *new(TokenBucketState), nil
}

// CosmosDBTokenBucketItem represents a document in CosmosDB.
type CosmosDBTokenBucketItem struct {
	ID           string           `json:"id"`
	PartitionKey string           `json:"partitionKey"`
	State        TokenBucketState `json:"state"`
	Version      int64            `json:"version"`
	TTL          int64            `json:"ttl,omitempty"`
}

// TokenBucketCosmosDB is a CosmosDB implementation of a TokenBucketStateBackend.
type TokenBucketCosmosDB struct {
	client        *azcosmos.ContainerClient
	partitionKey  string
	id            string
	ttl           time.Duration
	raceCheck     bool
	latestVersion int64
}

// NewTokenBucketCosmosDB creates a new TokenBucketCosmosDB instance.
// PartitionKey is the key used to store all the implementation in CosmosDB.
// TTL is the TTL of the stored item.
//
// If raceCheck is true and the item in CosmosDB is modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewTokenBucketCosmosDB(client *azcosmos.ContainerClient, partitionKey string, ttl time.Duration, raceCheck bool) *TokenBucketCosmosDB {
	_ = "STUB: not implemented"
	return nil
}

func (t *TokenBucketCosmosDB) State(ctx context.Context) (TokenBucketState, error) {
	_ = "STUB: not implemented"
	return *new(TokenBucketState), nil
}

func (t *TokenBucketCosmosDB) SetState(ctx context.Context, state TokenBucketState) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TokenBucketCosmosDB) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
