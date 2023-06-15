package main

import (
	"fmt"
	"log"

	"github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	client, err := api.DefaultRESTClient()
	if err != nil {
		log.Fatal(err)
	}
	response := struct {
		Login string
	}{}
	err = client.Get("user", &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Login)
}
