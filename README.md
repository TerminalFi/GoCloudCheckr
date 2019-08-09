# GoCloudCheckr

#### Not Completed

Golang API Library for CloudCheckr.

```
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/theseceng/gocloudcheckr/cloudcheckr"
)

func main() {

  // Create New API Client
	client := cloudcheckr.NewEnvClient(nil, "")
  
  // Set debugging true. Will generate corresponding curl command for terminal testing
	client.SetDebug(true)
  
  // Get Route Tables 
	details, err := client.GetRouteTables(context.Background(), nil)

	if err != nil {
		log.Panic(err)
	} else {
		data, _ := json.MarshalIndent(&details, "", "  ")
		fmt.Println(string(data))
	}
}
```
