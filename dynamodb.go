package limiters

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDBTableProperties are supplied to DynamoDB limiter backends.
// This struct informs the backend what the name of the table is and what the names of the key fields are.
type DynamoDBTableProperties struct {
	// TableName is the name of the table.
	TableName string
	// PartitionKeyName is the name of the PartitionKey attribute.
	PartitionKeyName string
	// SortKeyName is the name of the SortKey attribute.
	SortKeyName string
	// SortKeyUsed indicates if a SortKey is present on the table.
	SortKeyUsed bool
	// TTLFieldName is the name of the attribute configured for TTL.
	TTLFieldName string
}

// LoadDynamoDBTableProperties fetches a table description with the supplied client and returns a DynamoDBTableProperties struct.
func LoadDynamoDBTableProperties(ctx context.Context, client *dynamodb.Client, tableName string) (DynamoDBTableProperties, error) {
	_ = "STUB: not implemented"
	return *new(DynamoDBTableProperties), nil
}

func loadTableData(table *types.TableDescription, ttl *types.TimeToLiveDescription) (DynamoDBTableProperties, error) {
	_ = "STUB: not implemented"
	return *new(DynamoDBTableProperties), nil
}

func loadTableKeys(data DynamoDBTableProperties, table *types.TableDescription) (DynamoDBTableProperties, error) {
	_ = "STUB: not implemented"
	return *new(DynamoDBTableProperties), nil
}

func populateTableTTL(data DynamoDBTableProperties, ttl *types.TimeToLiveDescription) DynamoDBTableProperties {
	_ = "STUB: not implemented"
	return *new(DynamoDBTableProperties)
}

func dynamoDBputItem(ctx context.Context, client *dynamodb.Client, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dynamoDBGetItem(ctx context.Context, client *dynamodb.Client, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
