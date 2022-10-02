package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v47/github"
)

type Scrape struct {
	Repos []Repo
}

func newScrape() Scrape {
	var scrape Scrape
	client := login()
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	for {
		repos, resp, err := client.Repositories.List(context.Background(), "kijimaD", opt)
		if err != nil {
			panic(err)
		}
		for _, r := range repos {
			scrape.Repos = append(scrape.Repos, newRepo(r))
		}
		if resp.NextPage == 0 {
			break
		}
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
