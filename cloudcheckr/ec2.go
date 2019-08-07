package cloudcheckr

import "context"

const (
	getResourcesEC2AddressDetails      = "get_resources_ec2_address_details"
	getResourcesEC2Details             = "get_resources_ec2_details"
	getResourcesEC2DetailsV2           = "get_resources_ec2_details_V2"
	getResourcesEC2DetailsV3           = "get_resources_ec2_details_V3"
	getResourcesEC2HostDetails         = "get_resources_ec2_hosts"
	getResourcesEC2LoadBalancerDetails = "get_resources_ec2_load_balancer_details"
	getResourcesEC2LoadBalancerSummary = "ec2_load_balancer_summary"
	// getResourcesAmiSummary = "get_resources_ami_summary"
)

type AddressDetails struct {
	Addresses     *[]Addresses `json:"Addresses"`
	DateOfResults string       `json:"DateOfResults"`
	HasNext       bool         `json:"HasNext"`
	NextToken     string       `json:"NextToken"`
}

type Addresses struct {
	AccountName             string `json:"AccountName"`
	AllocationID            string `json:"AllocationId"`
	AssociationID           string `json:"AssociationId"`
	Domain                  string `json:"Domain"`
	InstanceID              string `json:"InstanceId"`
	NetworkInterfaceID      string `json:"NetworkInterfaceId"`
	NetworkInterfaceOwnerID string `json:"NetworkInterfaceOwnerId"`
	PublicIP                string `json:"PublicIp"`
	RegionDisplayName       string `json:"RegionDisplayName"`
	Status                  string `json:"Status"`
}

type Ec2Details struct {
	Count         int             `json:"Count"`
	DateOfResults string          `json:"DateOfResults"`
	Ec2Instances  *[]Ec2Instances `json:"Ec2Instances"`
	HasNext       bool            `json:"HasNext"`
	NextToken     string          `json:"NextToken"`
}

type Ec2Instances struct {
	AMI                           string            `json:"AMI"`
	AccountName                   string            `json:"AccountName"`
	AttachmentCount               int               `json:"AttachmentCount"`
	AvailabilityZone              string            `json:"AvailabilityZone"`
	AvgCpuforLast30Days           float64           `json:num_float"AvgCpuforLast30Days"`
	AvgCpuforLast7Days            float64           `json:num_float"AvgCpuforLast7Days"`
	AvgCpuforLast90Days           float64           `json:num_float"AvgCpuforLast90Days"`
	AvgNetworkInLast30Days        float64           `json:num_float"AvgNetworkInLast30Days"`
	AvgNetworkOutLast30Days       float64           `json:num_float"AvgNetworkOutLast30Days"`
	Cost                          float64           `json:num_float"Cost"`
	EbsCost                       float64           `json:num_float"EbsCost"`
	Ec2Cost                       float64           `json:num_float"Ec2Cost"`
	HighCPUPercent                int               `json:"HighCpuPercent"`
	HoursCPUUtilAbove80           int               `json:"HoursCpuUtilAbove80"`
	HoursCPUUtilBelow20           int               `json:"HoursCpuUtilBelow20"`
	HoursCPUUtilBelow40           int               `json:"HoursCpuUtilBelow40"`
	HoursCPUUtilBelow60           int               `json:"HoursCpuUtilBelow60"`
	HoursCPUUtilBelow80           int               `json:"HoursCpuUtilBelow80"`
	HoursHighCPULast30Days        int               `json:"HoursHighCpuLast30Days"`
	HoursHighCPULast7Days         int               `json:"HoursHighCpuLast7Days"`
	HoursHighCPULast90Days        int               `json:"HoursHighCpuLast90Days"`
	HoursLowCPULast30Days         int               `json:"HoursLowCpuLast30Days"`
	HoursLowCPULast7Days          int               `json:"HoursLowCpuLast7Days"`
	HoursLowCPULast90Days         int               `json:"HoursLowCpuLast90Days"`
	HoursRunningLast30Days        int               `json:"HoursRunningLast30Days"`
	HoursRunningLast7Days         int               `json:"HoursRunningLast7Days"`
	HoursRunningLast90Days        int               `json:"HoursRunningLast90Days"`
	Instance                      string            `json:"Instance"`
	InstanceID                    string            `json:"InstanceId"`
	InstanceName                  string            `json:"InstanceName"`
	InstanceType                  string            `json:"InstanceType"`
	LaunchTime                    string            `json:"LaunchTime"`
	LowCPUPercent                 int               `json:"LowCpuPercent"`
	MinimumCPUUtilization         float64           `json:num_float"MinimumCpuUtilization"`
	MinimumCPUUtilizationDateTime string            `json:"MinimumCpuUtilizationDateTime"`
	MonitoringState               string            `json:"MonitoringState"`
	PeakCPUUtilization            float64           `json:num_float"PeakCpuUtilization"`
	PeakCPUUtilizationDateTime    string            `json:"PeakCpuUtilizationDateTime"`
	Platform                      string            `json:"Platform"`
	PricingPlatform               string            `json:"PricingPlatform"`
	PricingType                   string            `json:"PricingType"`
	PrivateDNSName                string            `json:"PrivateDnsName"`
	PrivateIPAddress              string            `json:"PrivateIpAddress"`
	PublicDNSName                 string            `json:"PublicDnsName"`
	PublicIPAddress               string            `json:"PublicIpAddress"`
	RAMDiskID                     interface{}       `json:"RamDiskId"`
	Region                        string            `json:"Region"`
	ResourceTags                  *[]ResourceTags   `json:"ResourceTags"`
	SecurityGroups                *[]SecurityGroups `json:"SecurityGroups"`
	SourceDestinationCheck        bool              `json:"SourceDestinationCheck"`
	StateReasonCode               string            `json:"StateReasonCode"`
	StateReasonMessage            string            `json:"StateReasonMessage"`
	StateTransitionMessage        string            `json:"StateTransitionMessage"`
	Status                        string            `json:"Status"`
	SubnetID                      string            `json:"SubnetId"`
	Tenancy                       string            `json:"Tenancy"`
	TotalVolumeSizeGiB            int               `json:"TotalVolumeSizeGiB"`
	VirtualizationType            string            `json:"VirtualizationType"`
	VpcID                         string            `json:"VpcId"`
}

type ResourceTags struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type SecurityGroups struct {
	GroupID   string `json:"GroupId"`
	GroupName string `json:"GroupName"`
}

type Ec2HostDetails struct {
	Ec2HostDetails *[]Ec2HostDetail `json:"Ec2HostDetails"`
}

type Ec2HostDetail struct {
	AutoPlacement           string `json:"AutoPlacement"`
	AvailabilityZoneDetails string `json:"AvailabilityZoneDetails"`
	AvailableVCPUs          string `json:"AvailableVCPUs"`
	HostCores               string `json:"HostCores"`
	HostID                  string `json:"HostId"`
	HostInstanceType        string `json:"HostInstanceType"`
	HostReservationID       string `json:"HostReservationId"`
	HostSockets             string `json:"HostSockets"`
	HostTotalVCPUs          string `json:"HostTotalVCPUs"`
	State                   string `json:"State"`
}

type Ec2LoadBalancerDetails struct {
	DateOfResults    string             `json:"DateOfResults"`
	Ec2LoadBalancers *[]Ec2LoadBalancer `json:"Ec2LoadBalancers"`
	HasNext          bool               `json:"HasNext"`
	NextToken        string             `json:"NextToken"`
}

type Ec2LoadBalancer struct {
	AccountName               string          `json:"AccountName"`
	AvailableZones            string          `json:"AvailableZones"`
	CanonicalHostedZoneName   string          `json:"CanonicalHostedZoneName"`
	CanonicalHostedZoneNameID string          `json:"CanonicalHostedZoneNameId"`
	CreatedTime               string          `json:"CreatedTime"`
	DNSName                   string          `json:"DnsName"`
	HealthCheck               *HealthCheck    `json:"HealthCheck"`
	ID                        int             `json:"Id"`
	Listeners                 *[]Listeners    `json:"Listeners"`
	LoadBalancerName          string          `json:"LoadBalancerName"`
	RegionName                string          `json:"RegionName"`
	ResourceTags              *[]ResourceTags `json:"ResourceTags"`
}

type HealthCheck struct {
	HealthyThreshold   int    `json:"HealthyThreshold"`
	Interval           int    `json:"Interval"`
	Target             string `json:"Target"`
	Timeout            int    `json:"Timeout"`
	UnhealthyThreshold int    `json:"UnhealthyThreshold"`
}

type Listeners struct {
	InstancePort     int         `json:"InstancePort"`
	InstanceProtocol string      `json:"InstanceProtocol"`
	LoadBalancerPort int         `json:"LoadBalancerPort"`
	Protocol         string      `json:"Protocol"`
	SslCertificateID interface{} `json:"SslCertificateId"`
}

type Ec2LoadBalancerSummary struct {
	DateOfResults                        string                                  `json:"DateOfResults"`
	LoadBalancersByAccount               *[]LoadBalancersByAccount               `json:"LoadBalancersByAccount"`
	LoadBalancersByAttachedInstanceCount *[]LoadBalancersByAttachedInstanceCount `json:"LoadBalancersByAttachedInstanceCount"`
	LoadBalancersByRegion                *[]LoadBalancersByRegion                `json:"LoadBalancersByRegion"`
	TotalLoadBalancers                   int                                     `json:"TotalLoadBalancers"`
	TotalUnusedLoadBalancers             int                                     `json:"TotalUnusedLoadBalancers"`
}

type LoadBalancersByAccount struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type LoadBalancersByAttachedInstanceCount struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type LoadBalancersByRegion struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

// GetEC2AddressDetails https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_address_details
func (c *Client) GetEC2AddressDetails(ctx context.Context) (*AddressDetails, error) {
	var addressDetails AddressDetails

	req, err := c.NewInventoryRequest("GET", getResourcesEC2AddressDetails, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &addressDetails); err != nil {
		return nil, err
	}

	return &addressDetails, err
}

// https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_address_summary
// GetEC2AddressSummary https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_details
func (c *Client) GetEC2AddressSummary(ctx context.Context) (string, string) {
	return "Error", "Not implemented"
}

// GetEC2Details https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_details
func (c *Client) GetEC2Details(ctx context.Context) (*Ec2Details, error) {
	var ec2Details Ec2Details

	req, err := c.NewInventoryRequest("GET", getResourcesEC2Details, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ec2Details); err != nil {
		return nil, err
	}

	return &ec2Details, err
}

// GetEC2DetailsV2 https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_details_v2
func (c *Client) GetEC2DetailsV2(ctx context.Context) (*Ec2Details, error) {
	var ec2Details Ec2Details

	req, err := c.NewInventoryRequest("GET", getResourcesEC2DetailsV2, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ec2Details); err != nil {
		return nil, err
	}

	return &ec2Details, err
}

// GetEC2DetailsV3 https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_details_v3
func (c *Client) GetEC2DetailsV3(ctx context.Context) (*Ec2Details, error) {
	var ec2Details Ec2Details

	req, err := c.NewInventoryRequest("GET", getResourcesEC2DetailsV3, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ec2Details); err != nil {
		return nil, err
	}

	return &ec2Details, err
}

// GetEC2HostDetails https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_hosts
func (c *Client) GetEC2HostDetails(ctx context.Context) (*Ec2HostDetails, error) {
	var ec2HostDetails Ec2HostDetails

	req, err := c.NewInventoryRequest("GET", getResourcesEC2HostDetails, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ec2HostDetails); err != nil {
		return nil, err
	}

	return &ec2HostDetails, err
}

// GetEC2LoadBalancerDetails https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_load_balancer_details
func (c *Client) GetEC2LoadBalancerDetails(ctx context.Context) (*Ec2LoadBalancerDetails, error) {
	var ec2LoadBalancerDetails Ec2LoadBalancerDetails

	req, err := c.NewInventoryRequest("GET", getResourcesEC2LoadBalancerDetails, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ec2LoadBalancerDetails); err != nil {
		return nil, err
	}

	return &ec2LoadBalancerDetails, err
}

// GetEC2LoadBalancerSummary https://success.cloudcheckr.com/article/0ji9zuf1b8-api-reference-guide-inventory#ec2_load_balancer_details
func (c *Client) GetEC2LoadBalancerSummary(ctx context.Context) (*Ec2LoadBalancerSummary, error) {
	var ec2LoadBalancerSummary Ec2LoadBalancerSummary

	req, err := c.NewInventoryRequest("GET", getResourcesEC2LoadBalancerSummary, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ec2LoadBalancerSummary); err != nil {
		return nil, err
	}

	return &ec2LoadBalancerSummary, err
}
