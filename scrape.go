package main

import (
	"fmt"
	"github.com/google/go-github/v47/github"
	"context"
)

type Repo struct {
	Name            string
	Description     string
	Language        string
	HTMLURL         string
	DefaultBranch   string
	ForksCount      int
	StargazersCount int
	CommitCount     int
}

func newRepo(res *github.Repository) Repo {
	// ヌルポ対策
	var desc string
	var lang string
	if res.Description == nil {
		desc = ""
	} else {
		desc = *res.Description
	}
	if res.Language == nil {
		lang = ""
	} else {
		lang = *res.Language
	}

	result := Repo{
		*res.Name,
		desc, // *r.Description,
		lang, // *r.Language,
		*res.HTMLURL,
		*res.DefaultBranch,
		*res.ForksCount,
		*res.StargazersCount,
		0, // *contributors[0].Total,
	}
	fmt.Println(result)
	return result
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
