package cloudcheckr

import (
	"context"
)

const (
	getResourcesAmiDetails = "get_resources_ami_details"
	getResourcesAmiSummary = "get_resources_ami_summary"
)

// Profile holds account's information
type Amis struct {
	Amis          *[]Ami `json:"Amis"`
	DateOfResults string
}

type Ami struct {
	Architecture         string  `json:"Architecture"`
	AwsAccountID         string  `json:"AwsAccountId"`
	Cost                 float64 `json:num_float"Cost"`
	Description          string  `json:"Description"`
	InstanceCount        int     `json:"InstanceCount"`
	IsOwnedByAccount     bool    `json:"IsOwnedByAccount"`
	KernelID             string  `json:"KernelId"`
	Name                 string  `json:"Name"`
	OwnerID              string  `json:"OwnerId"`
	Platform             string  `json:"Platform"`
	Region               string  `json:"Region"`
	RunningInstanceCount int     `json:"RunningInstanceCount"`
	State                string  `json:"State"`
	StoppedInstanceCount int     `json:"StoppedInstanceCount"`
	StorageUsedBytes     int     `json:"StorageUsedBytes"`
	Type                 string  `json:"Type"`
	VirtualizationType   string  `json:"VirtualizationType"`
	Visibility           string  `json:"Visibility"`
	DateOfResults        string  `json:"DateOfResults"`
}

type AmiSummary struct {
	AmisByAccount    *[]AmisByAccount    `json:"AmisByAccount"`
	AmisByOwner      *[]AmisByOwner      `json:"AmisByOwner"`
	AmisByPlatform   *[]AmisByPlatform   `json:"AmisByPlatform"`
	AmisByRegion     *[]AmisByRegion     `json:"AmisByRegion"`
	AmisByVisibility *[]AmisByVisibility `json:"AmisByVisibility"`
	Cost             float64             `json:num_float"Cost"`
	DateOfResults    string              `json:"DateOfResults"`
	PaidAmiCount     int                 `json:"PaidAmiCount"`
	PrivateAmiCount  int                 `json:"PrivateAmiCount"`
	PublicAmiCount   int                 `json:"PublicAmiCount"`
}

type AmisByAccount struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type AmisByOwner struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type AmisByPlatform struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type AmisByRegion struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

type AmisByVisibility struct {
	Count int    `json:"Count"`
	Name  string `json:"Name"`
}

// GetAmiDetails returns information from CloudCheckr get_resources_ami_details API
func (c *Client) GetAmiDetails(ctx context.Context, parameters *Parameters) (*Amis, error) {
	var amis Amis

	req, err := c.NewInventoryRequest("GET", getResourcesAmiDetails, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &amis); err != nil {
		return nil, err
	}

	return &amis, err
}

// GetAmiSummary returns information from CloudCheckr get_resources_ami_details API
func (c *Client) GetAmiSummary(ctx context.Context, parameters *Parameters) (*AmiSummary, error) {
	var amisummary AmiSummary

	req, err := c.NewInventoryRequest("GET", getResourcesAmiSummary, parameters, nil)
	if err != nil {
		return nil, err
	}

	if err := c.Do(ctx, req, &amisummary); err != nil {
		return nil, err
	}

	return &amisummary, err
}
