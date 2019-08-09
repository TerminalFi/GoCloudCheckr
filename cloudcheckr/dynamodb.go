package cloudcheckr

import "context"

const (
	getResourcesDynamoDBDetails = "get_resources_dynamodb_details"
	getResourcesDynamoDBSummary = "get_resources_dynamodb_summary"
)

type DynamoDBDetails struct {
	DynamoDBDetails *[]DynamoDBDetail `json:"DynamoDbDetails"`
	DateOfResults   string            `json:"DateOfResults"`
	HasNext         bool              `json:"HasNext"`
	NextToken       interface{}       `json:"NextToken"`
}

type DynamoDBDetail struct {
	TableName                              string  `json:"TableName"`
	TableStatus                            string  `json:"TableStatus"`
	TableSizeBytes                         int     `json:"TableSizeBytes"`
	ItemCount                              int     `json:"ItemCount"`
	KeySchemaAttribute                     string  `json:"KeySchemaAttribute"`
	KeySchemaKeyType                       string  `json:"KeySchemaKeyType"`
	LocalSecondaryIndex                    int     `json:"LocalSecondaryIndex"`
	CreateTime                             string  `json:"CreateTime"`
	ProvisionedThroughputRead              int     `json:"ProvisionedThroughputRead"`
	ProvisionedThroughputWrite             int     `json:"ProvisionedThroughputWrite"`
	ProvisionedThroughputLastIncrease      string  `json:"ProvisionedThroughputLastIncrease"`
	ProvisionedThroughputLastDecrease      string  `json:"ProvisionedThroughputLastDecrease"`
	ProvisionedThroughputNumDecreasesToday int     `json:"ProvisionedThroughputNumDecreasesToday"`
	AverageConsumedReadsToday              int     `json:"AverageConsumedReadsToday"`
	AverageConsumedReadsLast2Days          int     `json:"AverageConsumedReadsLast2Days"`
	AverageConsumedReadsLast7Days          int     `json:"AverageConsumedReadsLast7Days"`
	AverageConsumedReadsLast30Days         int     `json:"AverageConsumedReadsLast30Days"`
	AverageConsumedWritesToday             int     `json:"AverageConsumedWritesToday"`
	AverageConsumedWritesLast2Days         int     `json:"AverageConsumedWritesLast2Days"`
	AverageConsumedWritesLast7Days         int     `json:"AverageConsumedWritesLast7Days"`
	AverageConsumedWritesLast30Days        int     `json:"AverageConsumedWritesLast30Days"`
	RegionName                             string  `json:"RegionName"`
	ProvisionedThroughputReadCost          float64 `json:num_float"ProvisionedThroughputReadCost"`
	ProvisionedThroughputWriteCost         float64 `json:num_float"ProvisionedThroughputWriteCost"`
	StorageCost                            float64 `json:num_float"StorageCost"`
	AwsAccountID                           int     `json:"AwsAccountId"`
}

type DynamoDBSummary struct {
	TotalTables                     int                `json:"TotalTables"`
	TotalActive                     int                `json:"TotalActive"`
	TotalStorage                    int                `json:"TotalStorage"`
	TotalItems                      int                `json:"TotalItems"`
	TotalProvisionedThroughputRead  int                `json:"TotalProvisionedThroughputRead"`
	TotalProvisionedThroughputWrite int                `json:"TotalProvisionedThroughputWrite"`
	TablesByRegion                  *[]TablesByRegion  `json:"TablesByRegion"`
	TablesByAccount                 *[]TablesByAccount `json:"TablesByAccount"`
	DateOfResults                   string             `json:"DateOfResults"`
}

type TablesByRegion struct {
	Name  string `json:"Name"`
	Count int    `json:"Count"`
}

type TablesByAccount struct {
	Name  string `json:"Name"`
	Count int    `json:"Count"`
}

// GetDynamoDBDetails https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#dynamo_db_details
func (c *Client) GetDynamoDBDetails(ctx context.Context, parameters *Parameters) (*DynamoDBDetails, error) {
	var dynamoDBDetails DynamoDBDetails

	req, err := c.NewInventoryRequest("GET", getResourcesDynamoDBDetails, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &dynamoDBDetails); err != nil {
		return nil, err
	}

	return &dynamoDBDetails, err
}

// GetDynamoDBSummary https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#dynamo_db_summary
func (c *Client) GetDynamoDBSummary(ctx context.Context, parameters *Parameters) (*DynamoDBSummary, error) {
	var dynamoDBSummary DynamoDBSummary
	req, err := c.NewInventoryRequest("GET", getResourcesDynamoDBSummary, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &dynamoDBSummary); err != nil {
		return nil, err
	}

	return &dynamoDBSummary, err
}
