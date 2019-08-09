package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/theseceng/gocloudcheckr/cloudcheckr"
)

func main() {

	// Create a new CloudCheckr Client for making API calls
	client := cloudcheckr.NewEnvClient(nil, "")

	// Enable debugging to log error messages and debug messages
	// Should probably seperate the two
	client.SetDebug(true)

	// Make a API call to get EC2 Address Details
	details, err := client.GetEC2AddressDetails(context.Background(), nil)

	if err != nil {
		client.Debug("Issue with API call", nil, err)
	} else {
		data, _ := json.MarshalIndent(&details, "", "  ")
		fmt.Println(string(data))
	}

}
