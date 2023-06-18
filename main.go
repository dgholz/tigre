package main

import (
	"fmt"
	"log"

	"github.com/cli/go-gh/v2/pkg/api"
)

type Tigre struct {
	client api.RESTClient
	user   string
}

func main() {
	client, err := api.DefaultRESTClient()
	if err != nil {
		log.Fatal(err)
	}
	tigre := Tigre{client: *client}

	tigre.set_user()
	fmt.Printf("%+v\n", tigre)
}

func (t *Tigre) set_user() {
	response := struct {
		Login string
	}{}
	err := t.client.Get("user", &response)

	if err != nil {
		log.Fatal(err)
	}
	t.user = response.Login
}
