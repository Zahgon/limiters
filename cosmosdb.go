package limiters

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type cosmosItem struct {
	Count        int64  `json:"Count"`
	PartitionKey string `json:"partitionKey"`
	ID           string `json:"id"`
	TTL          int32  `json:"ttl"`
}

// incrementCosmosItemRMW increments the count of a cosmosItem in a read-modify-write loop.
// It uses optimistic concurrency control (ETags) to ensure consistency.
func incrementCosmosItemRMW(ctx context.Context, client *azcosmos.ContainerClient, partitionKey, id string, ttl time.Duration) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
