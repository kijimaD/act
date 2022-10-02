package main

import (
	"context"
)

func main() {
	client := login()
	repos, _, _ := client.Repositories.List(context.Background(), "kijimaD", nil)
	for _, r := range repos {
		newRepo(r)
	}
}
