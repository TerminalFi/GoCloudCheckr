package cloudcheckr

import "context"

const (
	getResourcesEBSDetails   = "get_resources_ebs_details"
	getResourcesEBSSnapshots = "get_resources_ebs_snapshots"
	getResourcesEBSSummary   = "get_resources_ebs_summary"
)

type EBSDetails struct {
	EbsVolumes    *[]EBSVolumes `json:"EbsVolumes"`
	DateOfResults string        `json:"DateOfResults"`
	HasNext       bool          `json:"HasNext"`
	NextToken     string        `json:"NextToken"`
}
type Region struct {
	Name        string `json:"Name"`
	Location    string `json:"Location"`
	City        string `json:"City"`
	DisplayName string `json:"DisplayName"`
}
type AvailabilityZone struct {
	Name                       string  `json:"Name"`
	Region                     *Region `json:"Region"`
	ReservedInstanceOfferingID string  `json:"ReservedInstanceOfferingId"`
}
type EBSEc2Instance struct {
	ID                    int              `json:"Id"`
	Name                  string           `json:"Name"`
	DisplayName           string           `json:"DisplayName"`
	InstanceID            string           `json:"InstanceId"`
	LaunchTime            string           `json:"LaunchTime"`
	AvailabilityZone      AvailabilityZone `json:"AvailabilityZone"`
	StateTransitionReason interface{}      `json:"StateTransitionReason"`
	InstanceStateName     string           `json:"InstanceStateName"`
	InstanceType          string           `json:"InstanceType"`
	AmiPlatform           string           `json:"AmiPlatform"`
	PricingPlatform       string           `json:"PricingPlatform"`
	DisplayNameProperty   string           `json:"DisplayNameProperty"`
	MonitoringState       string           `json:"MonitoringState"`
	PrivateDNSName        string           `json:"PrivateDnsName"`
	PrivateIPAddress      string           `json:"PrivateIpAddress"`
	PublicIPAddress       interface{}      `json:"PublicIpAddress"`
	RootDeviceType        string           `json:"RootDeviceType"`
	PublicDNSName         interface{}      `json:"PublicDnsName"`
	RamdiskID             interface{}      `json:"RamdiskId"`
	ReservedInstanceIds   []interface{}    `json:"ReservedInstanceIds"`
	SourceDestCheck       bool             `json:"SourceDestCheck"`
	StateReasonCode       interface{}      `json:"StateReasonCode"`
	StateReasonMessage    interface{}      `json:"StateReasonMessage"`
	SubnetID              string           `json:"SubnetId"`
	VirtualizationType    string           `json:"VirtualizationType"`
	VpcID                 string           `json:"VpcId"`
	Tenancy               int              `json:"Tenancy"`
	HostID                interface{}      `json:"HostId"`
	SystemStatus          string           `json:"SystemStatus"`
	InstanceStatus        string           `json:"InstanceStatus"`
	PricingType           int              `json:"PricingType"`
	ImageID               string           `json:"ImageId"`
	PredictedHoursUsed    int              `json:"PredictedHoursUsed"`
	NetworkInterfaces     []interface{}    `json:"NetworkInterfaces"`
	AccountID             string           `json:"AccountId"`
	Identifier            string           `json:"Identifier"`
}

// type Attachments struct {
// 	AttachTime          string          `json:"AttachTime"`
// 	Device              string          `json:"Device"`
// 	DeleteOnTermination bool            `json:"DeleteOnTermination"`
// 	Status              string          `json:"Status"`
// 	InstanceID          string          `json:"InstanceId"`
// 	Ec2Instance         *EBSEc2Instance `json:"Ec2Instance"`
// }
type Costs struct {
	VolumeType  string      `json:"VolumeType"`
	IOPS        interface{} `json:"IOPS"`
	MonthlyCost float64     `json:"MonthlyCost"`
	CurrentCost bool        `json:"CurrentCost"`
}
type EBSVolumes struct {
	ID                   int             `json:"Id"`
	VolumeID             string          `json:"VolumeId"`
	VolumeStatus         string          `json:"VolumeStatus"`
	Cost                 float64         `json:"Cost"`
	StorageCost          float64         `json:"StorageCost"`
	PIOPSCost            float64         `json:"PIOPSCost"`
	Identifier           string          `json:"Identifier"`
	Encrypted            bool            `json:"Encrypted"`
	KmsKeyID             interface{}     `json:"KmsKeyId"`
	SizeGiB              int             `json:"SizeGiB"`
	SnapshotID           string          `json:"SnapshotId"`
	SnapshotCount        int             `json:"SnapshotCount"`
	SnapshotSizeGiB      int             `json:"SnapshotSizeGiB"`
	SnapshotCost         float64         `json:"SnapshotCost"`
	Created              string          `json:"Created"`
	Type                 string          `json:"Type"`
	Status               string          `json:"Status"`
	Region               string          `json:"Region"`
	AvailabilityZone     string          `json:"AvailabilityZone"`
	AttachmentToInstance string          `json:"AttachmentToInstance"`
	IOPS                 int             `json:"IOPS"`
	ResourceTags         *[]ResourceTags `json:"ResourceTags"`
	Snapshots            []interface{}   `json:"Snapshots"`
	Attachments          *[]Attachments  `json:"Attachments"`
	Costs                []Costs         `json:"Costs"`
	AccountName          string          `json:"AccountName"`
}

type EBSSnapshots struct {
	EbsSnapshots  []EBSSnapshot `json:"EbsSnapshots"`
	DateOfResults string        `json:"DateOfResults"`
}

type EBSSnapshot struct {
	SnapshotID   string          `json:"SnapshotId"`
	Region       string          `json:"Region"`
	VolumeID     string          `json:"VolumeId"`
	StartTime    string          `json:"StartTime"`
	Status       string          `json:"Status"`
	OwnerID      string          `json:"OwnerId"`
	Progress     string          `json:"Progress"`
	VolumeSize   string          `json:"VolumeSize"`
	Description  string          `json:"Description"`
	ResourceTags *[]ResourceTags `json:"ResourceTags"`
	AwsAccountID string          `json:"AwsAccountId"`
}

type EBSSummary struct {
	VolumeCount          int                 `json:"VolumeCount"`
	StorageUsedGiB       int                 `json:"StorageUsedGiB"`
	CostPerMonth         float64             `json:"CostPerMonth"`
	AvailableVolumeCount int                 `json:"AvailableVolumeCount"`
	InUseVolumeCount     int                 `json:"InUseVolumeCount"`
	VolumesByRegion      *[]VolumesByRegion  `json:"VolumesByRegion"`
	CostByRegion         *[]CostByRegion     `json:"CostByRegion"`
	VolumesByType        *[]VolumesByType    `json:"VolumesByType"`
	VolumesByAccount     *[]VolumesByAccount `json:"VolumesByAccount"`
	DateOfResults        string              `json:"DateOfResults"`
}
type VolumesByRegion struct {
	Name  string `json:"Name"`
	Count int    `json:"Count"`
}
type CostByRegion struct {
	Name string  `json:"Name"`
	Cost float64 `json:"Cost"`
}
type VolumesByType struct {
	Name  string `json:"Name"`
	Count int    `json:"Count"`
}
type VolumesByAccount struct {
	Name  string `json:"Name"`
	Count int    `json:"Count"`
}

func (c *Client) GetEBSDetails(ctx context.Context, parameters *Parameters) (*EBSDetails, error) {
	var ebsDetails EBSDetails

	req, err := c.NewInventoryRequest("GET", getResourcesEBSDetails, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ebsDetails); err != nil {
		return nil, err
	}

	return &ebsDetails, err
}

func (c *Client) GetEBSSnapshots(ctx context.Context, parameters *Parameters) (*EBSSnapshots, error) {
	var ebsSnapshots EBSSnapshots

	req, err := c.NewInventoryRequest("GET", getResourcesEBSSnapshots, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ebsSnapshots); err != nil {
		return nil, err
	}

	return &ebsSnapshots, err
}

func (c *Client) GetEBSSummary(ctx context.Context, parameters *Parameters) (*EBSSummary, error) {
	var ebsSummary EBSSummary

	req, err := c.NewInventoryRequest("GET", getResourcesEBSSummary, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &ebsSummary); err != nil {
		return nil, err
	}

	return &ebsSummary, err
}
