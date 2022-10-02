package utils

import (
	"context"
	"golang.org/x/oauth2"
	"os"
	"net/http"
)

func Login() *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	return tc
}
