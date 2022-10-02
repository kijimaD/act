package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v47/github"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
	"os"
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
			// commitCount := commitCount(".emacs.d")
			commitCount := 1
			scrape.Repos = append(scrape.Repos, newRepo(r, commitCount))
		}
		if resp.NextPage == 0 {
			break
		}
	}
	fmt.Println(commitCount(".emacs.d"))
	return scrape
}

func commitCount(reponame string) int {
	// GitHub REST APIでリポジトリの総コミット数を知る方法がなかったので、GraphQLを使っている
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	client := graphql.NewClient("https://api.github.com/graphql", tc)

	query := "{repository(owner:\"kijimaD\", name:\"dotfiles\") {object(expression:\"main\") {... on Commit {history {totalCount}}}}}"

	var res struct {
		Repository struct {
			Object struct {
				Commit struct {
					History struct {
						TotalCount int
					}
				} `graphql:"... on Commit"`
			} `graphql:"object(expression:\"master\")"`
		} `graphql:"repository(owner: \"kijimaD\", name: \".emacs.d\")"`
	}

	err := client.Exec(context.Background(), query, &res, nil)
	if err != nil {
		panic(err)
	}
	count := res.Repository.Object.Commit.History.TotalCount
	return count
}
