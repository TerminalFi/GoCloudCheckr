package cloudcheckr

import "context"

const (
	getResourcesCloudSearchDetails = "get_resources_cloudsearch_details"
	getResourcesCloudSearchSummary = "get_resources_cloudsearch_summary"
)

type CloudSearchDetails struct {
	CloudSearchDetails *[]CloudSearchDetail `json:"CloudSearchDetails"`
	DateOfResults      string               `json:"DateOfResults"`
	HasNext            bool                 `json:"HasNext"`
	NextToken          interface{}          `json:"NextToken"`
}

type CloudSearchDetail struct {
	AccountName                 interface{}                   `json:"AccountName"`
	AwsAccountID                interface{}                   `json:"AwsAccountId"`
	CloudSearchIndexFields      *[]CloudSearchIndexFields     `json:"CloudSearchIndexFields"`
	CloudSearchRankExpressions  *[]CloudSearchRankExpressions `json:"CloudSearchRankExpressions"`
	DefaultSearchFieldOption    interface{}                   `json:"DefaultSearchFieldOption"`
	DefaultSearchFieldStatus    interface{}                   `json:"DefaultSearchFieldStatus"`
	DomainID                    string                        `json:"DomainId"`
	DomainName                  string                        `json:"DomainName"`
	FieldNames                  interface{}                   `json:"FieldNames"`
	IsProcessing                bool                          `json:"IsProcessing"`
	RegionName                  string                        `json:"RegionName"`
	SearchInstanceType          string                        `json:"SearchInstanceType"`
	ServiceAccessPoliciesOption string                        `json:"ServiceAccessPoliciesOption"`
	ServiceAccessPoliciesStatus string                        `json:"ServiceAccessPoliciesStatus"`
	StemmingOptionsOption       string                        `json:"StemmingOptionsOption"`
	StemmingOptionsStatus       string                        `json:"StemmingOptionsStatus"`
	StopwordOptionsOption       []string                      `json:"StopwordOptionsOption"`
	StopwordOptionsStatus       string                        `json:"StopwordOptionsStatus"`
	SynonymOptionsOption        string                        `json:"SynonymOptionsOption"`
	SynonymOptionsStatus        string                        `json:"SynonymOptionsStatus"`
}

type CloudSearchIndexFields struct {
	CreationDate   string `json:"CreationDate"`
	IndexFieldName string `json:"IndexFieldName"`
	IndexFieldType string `json:"IndexFieldType"`
	State          string `json:"State"`
	UpdateDate     string `json:"UpdateDate"`
	UpdateVersion  int    `json:"UpdateVersion"`
}

type CloudSearchRankExpressions struct {
	CreationDate  string `json:"CreationDate"`
	RankName      string `json:"RankName"`
	State         string `json:"State"`
	UpdateDate    string `json:"UpdateDate"`
	UpdateVersion int    `json:"UpdateVersion"`
}

type CloudSearchSummary struct {
	ActiveOptionsCount  int                 `json:"ActiveOptionsCount"`
	DateOfResults       string              `json:"DateOfResults"`
	DomainCount         int                 `json:"DomainCount"`
	DomainsByAccount    *[]DomainsByAccount `json:"DomainsByAccount"`
	DomainsByRegion     *[]DomainsByRegion  `json:"DomainsByRegion"`
	IndexFeildCount     int                 `json:"IndexFeildCount"`
	RankExpressionCount int                 `json:"RankExpressionCount"`
}

type DomainsByRegion struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type DomainsByAccount struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

// GetCloudSearchDetails https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#cloud_formation_details
func (c *Client) GetCloudSearchDetails(ctx context.Context, parameters *Parameters) (*CloudSearchDetails, error) {
	var cloudSearchDetails CloudSearchDetails

	req, err := c.NewInventoryRequest("GET", getResourcesCloudSearchDetails, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &cloudSearchDetails); err != nil {
		return nil, err
	}

	return &cloudSearchDetails, err
}

// GetCloudSearchSummary https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#cloud_formation_details
func (c *Client) GetCloudSearchSummary(ctx context.Context, parameters *Parameters) (*CloudSearchSummary, error) {
	var cloudSearchSummary CloudSearchSummary

	req, err := c.NewInventoryRequest("GET", getResourcesCloudSearchSummary, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &cloudSearchSummary); err != nil {
		return nil, err
	}

	return &cloudSearchSummary, err
}
