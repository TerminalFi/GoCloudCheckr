package cloudcheckr

import "context"

const (
	getResourcesRouteTables = "get_resources_route_tables"
)

type RouteTables struct {
	VPCRouteTableDetails *[]VPCRouteTableDetails `json:"VPCRouteTableDetails"`
}
type VpcRouteTablesAssociationDetails struct {
	Associations string `json:"Associations"`
}
type EC2ResourceTagDetails struct {
	ResourceTags string `json:"ResourceTags"`
}
type VPCRouteTableDetails struct {
	ID                               string                              `json:"Id"`
	RouteTableID                     string                              `json:"RouteTableId"`
	Main                             string                              `json:"Main"`
	VpcRouteTablesAssociationDetails *[]VpcRouteTablesAssociationDetails `json:"VpcRouteTablesAssociationDetails"`
	EC2ResourceTagDetails            *[]EC2ResourceTagDetails            `json:"EC2ResourceTagDetails"`
	RegionName                       string                              `json:"RegionName"`
	AccountName                      string                              `json:"AccountName"`
	VPCID                            string                              `json:"VPCId"`
}

func (c *Client) GetRouteTables(ctx context.Context, parameters *Parameters) (*RouteTables, error) {
	var routeTables RouteTables

	req, err := c.NewInventoryRequest("GET", getResourcesRouteTables, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &routeTables); err != nil {
		return nil, err
	}

	return &routeTables, err
}
