package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/theseceng/gocloudcheckr/cloudcheckr"
)

func main() {

	client := cloudcheckr.NewEnvClient(nil, "")
	client.SetDebug(true)
	details, err := client.GetEC2AddressDetails(context.Background(), nil)

	if err != nil {
		log.Panic(err)
	} else {
		data, _ := json.MarshalIndent(&details, "", "  ")
		fmt.Println(string(data))
	}

}
