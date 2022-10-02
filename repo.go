package main

import (
	"github.com/google/go-github/v47/github"
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
		desc,
		lang,
		*res.HTMLURL,
		*res.DefaultBranch,
		*res.ForksCount,
		*res.StargazersCount,
		0, // *contributors[0].Total,
	}
	return result
}

func (r *Repo) headers() []string {
	return []string{"Name", "Description", "Language", "URL", "Forks", "Star", "Commit"}
}
