package cloudcheckr

import "context"

const (
	getResourcesCustomerGateways = "get_resources_customer_gateways"
)

type CustomerGateway struct {
	Account           string             `json:"Account"`
	BgpAsn            string             `json:"BgpAsn"`
	CustomerGatewayID string             `json:"CustomerGatewayId"`
	ID                int                `json:"Id"`
	IPAddress         string             `json:"IpAddress"`
	Region            string             `json:"Region"`
	VPCResourceTags   *[]VPCResourceTags `json:"ResourceTags"`
	State             string             `json:"State"`
	Type              string             `json:"Type"`
}

type VPCResourceTags struct {
	TagKey   string `json:"TagKey"`
	TagValue string `json:"TagValue"`
}

// GetVPCCustomerGateways https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#cloud_formation_details
func (c *Client) GetCustomerGateways(ctx context.Context, parameters *Parameters) (*[]CustomerGateway, error) {
	var customerGateway []CustomerGateway

	req, err := c.NewInventoryRequest("GET", getResourcesCustomerGateways, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &customerGateway); err != nil {
		return nil, err
	}

	return &customerGateway, err
}
