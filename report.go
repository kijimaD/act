package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"strconv"
)

type Report struct {
	in     []Repo
	out    string
	config Config
}

func newReport(in Scrape, config Config) *Report {
	return &Report{
		in:     in.Repos,
		out:    "",
		config: config,
	}
}

func (r *Report) execute() {
	// 出力分岐を別の関数にしたいが、fileが空白になる
	var output io.Writer

	switch r.config.OutType {
	case "stdout":
		output = os.Stdout
	case "file":
		file, err := os.Create("README.md")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		output = file
	default:
		panic("output type error. stdout | file")
	}

	table := tablewriter.NewWriter(output)
	// table := tablewriter.NewWriter(r.selectOutput())
	table.SetHeader(r.headers())
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)

	for _, repo := range r.in {
		table.Append(r.content(repo))
	}

	table.Render()
}

func (r *Report) headers() []string {
	return []string{"Name", "Description", "Language", "Forks", "Star", "Commit"}
}

func (r *Report) content(repo Repo) []string {
	return []string{repo.MdLink(), repo.Description, repo.Language, strconv.Itoa(repo.ForksCount), strconv.Itoa(repo.StargazersCount), strconv.Itoa(repo.CommitCount)}
}
