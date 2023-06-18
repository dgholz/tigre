package main

import (
	"fmt"
	"log"

	"github.com/cli/go-gh/v2/pkg/api"
)

type Tigre struct {
	client api.RESTClient
}

func main() {
	client, err := api.DefaultRESTClient()
	if err != nil {
		log.Fatal(err)
	}
	tigre := Tigre{client: *client}

	response := struct {
		Login string
	}{}
	err = tigre.client.Get("user", &response)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Login)
}
