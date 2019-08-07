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
	// client.SetDebug(true)
	details, err := client.GetEC2Details(context.Background())

	if err != nil {
		log.Panic(err)
	} else {
		data, _ := json.MarshalIndent(&details, "", "  ")
		fmt.Println(string(data))
	}
	// amis, err := client.GetAmiDetails(context.Background())

	// if err != nil {
	// 	log.Panic(err)
	// } else {
	// 	data, _ := json.MarshalIndent(&amis, "", "  ")
	// 	log.Println(string(data))
	// }

	// amisummary, err := client.GetAmiSummary(context.Background())
	// if err != nil {
	// 	log.Panic(err)
	// } else {
	// 	data, _ := json.MarshalIndent(&amisummary, "", "  ")
	// 	log.Println(string(data))
	// }

}
