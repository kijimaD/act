package main

import(
	"github.com/google/go-github/v47/github"
	"fmt"
	"context"
)

type Stats struct {
	Name string
	Description string
	Language string
	HTMLURL string
	DefaultBranch string
	ForksCount int
	StargazersCount int
}

func main() {
	client := github.NewClient(nil)

	// list all organizations for user
	orgs, _, err := client.Organizations.List(context.Background(), "kijimaD", nil)
	if err != nil{
		panic("error")
	}
	fmt.Println(orgs)

	repos, _ , _ := client.Repositories.List(context.Background(), "kijimaD", nil)
	fmt.Println(*repos[0])
	fmt.Println(*repos[0].Name)
	fmt.Println(*repos[0].Description)
	fmt.Println(*repos[0].Language)
	fmt.Println(*repos[0].HTMLURL)
	fmt.Println(*repos[0].DefaultBranch)
	fmt.Println(*repos[0].ForksCount)
	fmt.Println(*repos[0].StargazersCount)

	stats, _, _ := client.Repositories.ListContributorsStats(context.Background(), "kijimaD", *repos[0].Name)
	fmt.Println(*stats[0].Total)
}
