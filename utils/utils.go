package utils

import (
	"context"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"os"
)

func Login() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
