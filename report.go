package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type Report struct {
	in  []Repo
	out string
}

func newReport(in Scrape) *Report {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "Language"})

	for _, repo := range in.Repos {
		table.Append([]string{repo.Name, repo.Description, repo.Language})
	}

	table.Render()

	return &Report{
		in:  in.Repos,
		out: "",
	}
}
