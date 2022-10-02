package main

import (
	"context"
	"fmt"
)

func main() {
	var scrape Scrape

	client := login()
	rs, _, _ := client.Repositories.List(context.Background(), "kijimaD", nil)
	for _, r := range rs {
		scrape.Repos = append(scrape.Repos, newRepo(r))
	}
	fmt.Println(scrape.Repos)
}
