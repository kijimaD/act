package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type Report struct {
	in  []Repo
	out string
}

func newReport(in Scrape) *Report {
	return &Report{
		in:  in.Repos,
		out: "",
	}
}

func (r *Report) execute() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(r.headers())

	for _, repo := range r.in {
		table.Append(r.content(repo))
	}

	table.Render()
}

func (r *Report) headers() []string {
	return []string{"Name", "Description", "Language", "URL", "Forks", "Star", "Commit"}
}

func (r *Report) content(repo Repo) []string {
	return []string{repo.Name, repo.Description, repo.Language, repo.HTMLURL, strconv.Itoa(repo.ForksCount), strconv.Itoa(repo.StargazersCount), strconv.Itoa(repo.CommitCount)}
}
