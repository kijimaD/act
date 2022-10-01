package main

import(
	"github.com/google/go-github/v47/github"
	"fmt"
	"context"
	"os"
	"golang.org/x/oauth2"
)

type Repo struct {
	Name string
	Description string
	Language string
	HTMLURL string
	DefaultBranch string
	ForksCount int
	StargazersCount int
	CommitCount int
}

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repos, _ , _ := client.Repositories.List(ctx, "kijimaD", nil)
	for _, r := range repos {
		contributors, _, err := client.Repositories.ListContributorsStats(context.Background(), "kijimaD", *r.Name)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(*r.Name))
		fmt.Println(*contributors[0].Total)
		repo := Repo{
			// string(*r.Name),
			"a",
			// string(*r.Description),
			"a",
			// string(*r.Language),
			"a",
			// *r.HTMLURL,
			"a",
			// *r.DefaultBranch,
			"a",
			// *r.ForksCount,
			0,
			// *r.StargazersCount,
			0,
			// *contributors[0].Total,
			0,
		}
		fmt.Println(repo)
	}

	// repo := Repo{
	// 	*repos[0].Name,
	// 	*repos[0].Description,
	// 	*repos[0].Language,
	// 	*repos[0].HTMLURL,
	// 	*repos[0].DefaultBranch,
	// 	*repos[0].ForksCount,
	// 	*repos[0].StargazersCount,
	// 	*contributors[0].Total,
	// }
}
