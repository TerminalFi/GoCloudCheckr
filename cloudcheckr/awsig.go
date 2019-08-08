package cloudcheckr

import "context"

const (
	getResourcesAWSInternetGateways = "get_resources_aws_internet_gateways"
)

type AWSInternetGateways struct {
	Account           string               `json:"Account"`
	Attachments       *[]Attachments       `json:"Attachments"`
	ID                int                  `json:"Id"`
	InternetGatewayID string               `json:"InternetGatewayId"`
	Region            string               `json:"Region"`
	ResourceTags      *[]AWSIGResourceTags `json:"ResourceTags"`
	State             string               `json:"State"`
}

type Attachments struct {
	State string `json:"State"`
	VpcID string `json:"VpcId"`
}

type AWSIGResourceTags struct {
	TagKey   string `json:"TagKey"`
	TagValue string `json:"TagValue"`
}

// GetAWSInternetGateways https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#aws_internet_gateways
func (c *Client) GetAWSInternetGateways(ctx context.Context, parameters *Parameters) (*AWSInternetGateways, error) {
	var awsInternetGateways AWSInternetGateways

	req, err := c.NewInventoryRequest("GET", getResourcesAWSInternetGateways, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &awsInternetGateways); err != nil {
		return nil, err
	}

	return &awsInternetGateways, err
}
