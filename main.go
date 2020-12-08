package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	// Login and authenticate
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GHCLI_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// List repos
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}
