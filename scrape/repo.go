package scrape

import (
	"fmt"
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

func NewRepo(res *github.Repository, commitCount int) Repo {
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
		commitCount,
	}
	return result
}

func (r *Repo) MdLink() string {
	return fmt.Sprintf("[%s](%s)", r.Name, r.HTMLURL)
}
