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

type Repo struct {
	Name string
}

func main() {
	client, err := api.DefaultRESTClient()
	if err != nil {
		log.Fatal(err)
	}
	tigre := Tigre{client: *client}

	tigre.set_user()
	fmt.Printf("%+v\n", tigre)
	repos := make(chan Repo)
	go tigre.user_repos(repos)
	for repo := range repos {
		fmt.Printf("%+v\n", repo)
	}
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

func (t Tigre) user_repos(c chan Repo) {
	var response []Repo
	err := t.client.Get(fmt.Sprintf("users/%s/repos", t.user), &response)

	if err != nil {
		log.Fatal(err)
	}
	for _, repo := range response {
		c <- repo
	}
	close(c)
}
