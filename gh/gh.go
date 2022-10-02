package gh

import (
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/v47/github"
	"os"
	"net/http"
)

type gh struct {
	Client *github.Client
}

func New() gh {
	client := github.NewClient(Login())

	return gh{
		Client: client,
	}
}

func Login() *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	return tc
}
