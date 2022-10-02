package main

import (
	"context"
	"fmt"
)

type Scrape struct {
	Repos []Repo
}

func newScrape() Scrape {
	var scrape Scrape
	client := login()
	repos, _, _ := client.Repositories.List(context.Background(), "kijimaD", nil)
	for _, r := range repos {
		scrape.Repos = append(scrape.Repos, newRepo(r))
	}
	return scrape
}

func contrib(reponame string) string {
	client := login()
	context := context.Background()
	// TODO: ContributionStatsは、キャッシュが未生成の場合少し待って再リトライしてあげないといけない
	contributors, _, err := client.Repositories.ListContributorsStats(context, "kijimaD", reponame)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	} else {
		fmt.Println(*contributors[0].Total)
	}

	return "a"
}
