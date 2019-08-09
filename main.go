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
<<<<<<< HEAD
	details, err := client.GetEC2AddressDetails(context.Background(), nil)
=======
	details, err := client.GetRouteTables(context.Background(), nil)
>>>>>>> cbeade4ab65e7e8f1a4400d399b2b8f653e2f389

	if err != nil {
		log.Panic(err)
	} else {
		data, _ := json.MarshalIndent(&details, "", "  ")
		fmt.Println(string(data))
	}
<<<<<<< HEAD

=======
>>>>>>> cbeade4ab65e7e8f1a4400d399b2b8f653e2f389
}
