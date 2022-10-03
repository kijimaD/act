package scrape

import (
	"context"
	"fmt"
	"github.com/google/go-github/v47/github"
	"github.com/hasura/go-graphql-client"
	"act/config"
	"act/gh"
)

type Scrape struct {
	Repos  []Repo
	config config.Config
}

func NewScrape(config config.Config) *Scrape {
	var scrape Scrape
	scrape.config = config
	return &scrape
}

func (s *Scrape) Execute() *Scrape {
	client := gh.New(s.config).Client
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	for {
		repos, resp, err := client.Repositories.List(context.Background(), s.config.UserId, opt)
		if err != nil {
			panic(err)
		}
		for _, r := range repos {
			commitCount := s.commitCount(*r.Name, *r.DefaultBranch)
			s.Repos = append(s.Repos, NewRepo(r, commitCount))
		}
		if resp.NextPage == 0 {
			break
		}
	}
	return s
}

func (s *Scrape) commitCount(reponame string, branch string) int {
	// GitHub REST APIでリポジトリの総コミット数を知る方法がなかったので、GraphQLを使っている
	client := graphql.NewClient("https://api.github.com/graphql", gh.Login())

	// 変数展開が必要なためクエリを文字列モードで実行する
	query := fmt.Sprintf("{repository(owner:\"%s\", name:\"%s\") {object(expression:\"%s\") {... on Commit {history {totalCount}}}}}", s.config.UserId, reponame, branch)

	var res struct {
		Repository struct {
			Object struct {
				Commit struct {
					History struct {
						TotalCount int
					}
				} `graphql:"... on Commit"`
			}
		}
	}

	err := client.Exec(context.Background(), query, &res, nil)
	if err != nil {
		panic(err)
	}
	count := res.Repository.Object.Commit.History.TotalCount
	return count
}
