package cloudcheckr

import "context"

const (
	getResourcesCloudFormationDetails = "get_resources_cloudformation_details"
	getResourcesCloudFormationSummary = "get_resources_cloudformation_summary"
)

type CloudFormationDetails struct {
	CloudFormationDetail *[]CloudFormationDetail `json:"CloudFormationDetails"`
	DateOfResults        string                  `json:"DateOfResults"`
	HasNext              bool                    `json:"HasNext"`
	NextToken            interface{}             `json:"NextToken"`
}

type CloudFormationDetail struct {
	AccountName      string       `json:"AccountName"`
	Capabilities     string       `json:"Capabilities"`
	Creation         string       `json:"Creation"`
	Description      string       `json:"Description"`
	DisableRollback  bool         `json:"DisableRollback"`
	LastUpdated      string       `json:"LastUpdated"`
	Name             string       `json:"Name"`
	NotificationARNs interface{}  `json:"NotificationARNs"`
	Outputs          string       `json:"Outputs"`
	Parameters       string       `json:"Parameters"`
	RegionName       string       `json:"RegionName"`
	Resources        *[]Resources `json:"Resources"`
	StackID          string       `json:"StackId"`
	Status           string       `json:"Status"`
	StatusReason     interface{}  `json:"StatusReason"`
	TimeoutInMinutes int          `json:"TimeoutInMinutes"`
}

type Resources struct {
	ID                   int         `json:"Id"`
	LastUpdated          string      `json:"LastUpdated"`
	LogicalResourceID    string      `json:"LogicalResourceId"`
	PhysicalResourceID   string      `json:"PhysicalResourceId"`
	ResourceStatus       string      `json:"ResourceStatus"`
	ResourceStatusReason interface{} `json:"ResourceStatusReason"`
	ResourceType         string      `json:"ResourceType"`
}

type CloudFormationSummary struct {
	DateOfResults   string             `json:"DateOfResults"`
	StacksByAccount *[]StacksByAccount `json:"StacksByAccount"`
	StacksByRegion  *[]StacksByRegion  `json:"StacksByRegion"`
	TotalResources  int                `json:"TotalResources"`
	TotalStacks     int                `json:"TotalStacks"`
}
type StacksByAccount struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type StacksByRegion struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

// GetCloudFormationDetails https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#cloud_formation_details
func (c *Client) GetCloudFormationDetails(ctx context.Context, parameters *Parameters) (*CloudFormationDetails, error) {
	var cloudFormationDetails CloudFormationDetails

	req, err := c.NewInventoryRequest("GET", getResourcesCloudFormationDetails, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &cloudFormationDetails); err != nil {
		return nil, err
	}

	return &cloudFormationDetails, err
}

// GetCloudFormationSummary https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#cloud_formation_summary
func (c *Client) GetCloudFormationSummary(ctx context.Context, parameters *Parameters) (*CloudFormationSummary, error) {
	var cloudFormationSummary CloudFormationSummary

	req, err := c.NewInventoryRequest("GET", getResourcesCloudFormationSummary, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &cloudFormationSummary); err != nil {
		return nil, err
	}

	return &cloudFormationSummary, err
}
