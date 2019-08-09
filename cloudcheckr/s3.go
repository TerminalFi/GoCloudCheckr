package cloudcheckr

import "context"

const (
	getResourcesS3Details = "get_resources_s3_details"
	getResourcesS3Summary = "get_resources_s3_summary"
)

type S3Buckets struct {
	S3Buckets     *[]S3Bucket `json:"S3Buckets"`
	DateOfResults string      `json:"DateOfResults"`
}

type S3Permissions struct {
	Grantee     string   `json:"Grantee"`
	Permissions []string `json:"Permissions"`
}

type ObjectTypesTop10 struct {
	Type             string  `json:"Type"`
	ObjectCount      int     `json:"ObjectCount"`
	PercentOfObjects float64 `json:"PercentOfObjects"`
	StorageUsedBytes int     `json:"StorageUsedBytes"`
	PercentOfStorage float64 `json:"PercentOfStorage"`
}

type ObjectsBySizeTop10 struct {
	Key          string `json:"Key"`
	SizeBytes    int    `json:"SizeBytes"`
	LastModified string `json:"LastModified"`
	StorageClass string `json:"StorageClass"`
}

type ObjectsByMostRecentlyModifiedTop10 struct {
	Key          string `json:"Key"`
	SizeBytes    int    `json:"SizeBytes"`
	LastModified string `json:"LastModified"`
	StorageClass string `json:"StorageClass"`
}

type ObjectsByLongestSinceModifiedTop10 struct {
	Key          string `json:"Key"`
	SizeBytes    int    `json:"SizeBytes"`
	LastModified string `json:"LastModified"`
	StorageClass string `json:"StorageClass"`
}

type S3Bucket struct {
	Name                               string                                `json:"Name"`
	Region                             string                                `json:"Region"`
	EstimatedCost                      float64                               `json:"EstimatedCost"`
	StorageUsedBytes                   int                                   `json:"StorageUsedBytes"`
	ObjectCount                        int                                   `json:"ObjectCount"`
	FolderCount                        int                                   `json:"FolderCount"`
	Created                            string                                `json:"Created"`
	AccessLogCount                     int                                   `json:"AccessLogCount"`
	AccessLogSizeBytes                 int                                   `json:"AccessLogSizeBytes"`
	WebsiteEnabled                     bool                                  `json:"WebsiteEnabled"`
	WebsiteIndexDocument               interface{}                           `json:"WebsiteIndexDocument"`
	WebsiteErrorDocument               interface{}                           `json:"WebsiteErrorDocument"`
	LoggingEnabled                     bool                                  `json:"LoggingEnabled"`
	LoggingTargetBucket                string                                `json:"LoggingTargetBucket"`
	LoggingPrefix                      interface{}                           `json:"LoggingPrefix"`
	NotificationEnabled                bool                                  `json:"NotificationEnabled"`
	IsPublic                           bool                                  `json:"IsPublic"`
	StandardCost                       float64                               `json:"StandardCost"`
	StandardObjectCount                int                                   `json:"StandardObjectCount"`
	StandardStorageUsedBytes           int                                   `json:"StandardStorageUsedBytes"`
	RRCost                             float64                               `json:"RRCost"`
	RRObjectCount                      int                                   `json:"RRObjectCount"`
	RRStorageUsedBytes                 int                                   `json:"RRStorageUsedBytes"`
	GlacierCost                        float64                               `json:"GlacierCost"`
	GlacierObjectCount                 int                                   `json:"GlacierObjectCount"`
	GlacierStorageUsedBytes            int                                   `json:"GlacierStorageUsedBytes"`
	AvgObjectSizeBytes                 int                                   `json:"AvgObjectSizeBytes"`
	AvgStandardObjectSizeBytes         int                                   `json:"AvgStandardObjectSizeBytes"`
	AvgRRObjectSizeBytes               int                                   `json:"AvgRRObjectSizeBytes"`
	AvgGlacierObjectSizeBytes          int                                   `json:"AvgGlacierObjectSizeBytes"`
	AvgCostPerObject                   float64                               `json:"AvgCostPerObject"`
	AvgCostPerStandardObject           float64                               `json:"AvgCostPerStandardObject"`
	AvgCostPerRRObject                 float64                               `json:"AvgCostPerRRObject"`
	AvgCostPerGlacierObject            float64                               `json:"AvgCostPerGlacierObject"`
	AvgTimeSinceLastModified           string                                `json:"AvgTimeSinceLastModified"`
	S3Permissions                      *[]S3Permissions                      `json:"S3Permissions"`
	S3Notifications                    []interface{}                         `json:"S3Notifications"`
	S3LifeCycleRules                   []interface{}                         `json:"S3LifeCycleRules"`
	ResourceTags                       *[]ResourceTags                       `json:"ResourceTags"`
	ObjectTypesTop10                   *[]ObjectTypesTop10                   `json:"ObjectTypesTop10"`
	ObjectsBySizeTop10                 *[]ObjectsBySizeTop10                 `json:"ObjectsBySizeTop10"`
	ObjectsByMostRecentlyModifiedTop10 *[]ObjectsByMostRecentlyModifiedTop10 `json:"ObjectsByMostRecentlyModifiedTop10"`
	ObjectsByLongestSinceModifiedTop10 *[]ObjectsByLongestSinceModifiedTop10 `json:"ObjectsByLongestSinceModifiedTop10"`
	AwsAccountID                       string                                `json:"AwsAccountId"`
}

type S3BucketsSummary struct {
	CostPerMonth                      float64                             `json:"CostPerMonth"`
	ExceededCloudCheckrS3Limit        bool                                `json:"ExceededCloudCheckrS3Limit"`
	StorageUsedBytes                  int64                               `json:"StorageUsedBytes"`
	BucketCount                       int                                 `json:"BucketCount"`
	BucketCountLoggingEnabled         int                                 `json:"BucketCountLoggingEnabled"`
	BucketCountSnsNotificationEnabled int                                 `json:"BucketCountSnsNotificationEnabled"`
	ObjectCount                       int                                 `json:"ObjectCount"`
	AccessLogCount                    int                                 `json:"AccessLogCount"`
	AccessLogSizeBytes                int                                 `json:"AccessLogSizeBytes"`
	BucketCountOpenWritePermissions   int                                 `json:"BucketCountOpenWritePermissions"`
	BucketCountOpenListPermissions    int                                 `json:"BucketCountOpenListPermissions"`
	FolderCount                       int                                 `json:"FolderCount"`
	StandardCost                      float64                             `json:"StandardCost"`
	StandardObjectCount               int                                 `json:"StandardObjectCount"`
	StandardStorageUsedBytes          int64                               `json:"StandardStorageUsedBytes"`
	RRCost                            float64                             `json:"RRCost"`
	RRObjectCount                     int                                 `json:"RRObjectCount"`
	RRStorageUsedBytes                int                                 `json:"RRStorageUsedBytes"`
	GlacierCost                       float64                             `json:"GlacierCost"`
	GlacierObjectCount                int                                 `json:"GlacierObjectCount"`
	GlacierStorageUsedBytes           int64                               `json:"GlacierStorageUsedBytes"`
	AvgObjectSizeBytes                int                                 `json:"AvgObjectSizeBytes"`
	AvgStandardObjectSizeBytes        int                                 `json:"AvgStandardObjectSizeBytes"`
	AvgRRObjectSizeBytes              int                                 `json:"AvgRRObjectSizeBytes"`
	AvgGlacierObjectSizeBytes         int                                 `json:"AvgGlacierObjectSizeBytes"`
	AvgCostPerObject                  float64                             `json:"AvgCostPerObject"`
	AvgCostPerStandardObject          float64                             `json:"AvgCostPerStandardObject"`
	AvgCostPerRRObject                float64                             `json:"AvgCostPerRRObject"`
	AvgCostPerGlacierObject           float64                             `json:"AvgCostPerGlacierObject"`
	AvgTimeSinceLastModified          string                              `json:"AvgTimeSinceLastModified"`
	EstimatedMonthlyCostByRegion      *[]EstimatedMonthlyCostByRegion     `json:"EstimatedMonthlyCostByRegion"`
	ObjectsByRegion                   *[]ObjectsByRegion                  `json:"ObjectsByRegion"`
	EstimatedMonthlyCostByBucketTop5  *[]EstimatedMonthlyCostByBucketTop5 `json:"EstimatedMonthlyCostByBucketTop5"`
	ObjectsByBucketTop5               *[]ObjectsByBucketTop5              `json:"ObjectsByBucketTop5"`
	BucketsByAccount                  *[]BucketsByAccount                 `json:"BucketsByAccount"`
	ObjectTypesTop10                  *[]ObjectTypesTop10                 `json:"ObjectTypesTop10"`
	DateOfResults                     string                              `json:"DateOfResults"`
}
type EstimatedMonthlyCostByRegion struct {
	Name             string  `json:"Name"`
	Cost             float64 `json:"Cost"`
	StorageUsedBytes int64   `json:"StorageUsedBytes"`
}
type ObjectsByRegion struct {
	Name        string `json:"Name"`
	BucketCount int    `json:"BucketCount"`
	ObjectCount int    `json:"ObjectCount"`
}
type EstimatedMonthlyCostByBucketTop5 struct {
	Name             string  `json:"Name"`
	Cost             float64 `json:"Cost"`
	StorageUsedBytes int64   `json:"StorageUsedBytes"`
}

type ObjectsByBucketTop5 struct {
	Name        string `json:"Name"`
	BucketCount int    `json:"BucketCount"`
}

type BucketsByAccount struct {
	Name  string `json:"Name"`
	Count int    `json:"Count"`
}

func (c *Client) GetS3Details(ctx context.Context, parameters *Parameters) (*S3Buckets, error) {
	var s3Buckets S3Buckets

	req, err := c.NewInventoryRequest("GET", getResourcesS3Details, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &s3Buckets); err != nil {
		return nil, err
	}

	return &s3Buckets, err
}

func (c *Client) GetS3Summary(ctx context.Context, parameters *Parameters) (*S3BucketsSummary, error) {
	var s3BucketsSummary S3BucketsSummary

	req, err := c.NewInventoryRequest("GET", getResourcesS3Summary, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &s3BucketsSummary); err != nil {
		return nil, err
	}

	return &s3BucketsSummary, err
}
