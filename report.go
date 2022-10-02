package main

import (
	"fmt"
)

type Report struct {
	in  []Repo
	out string
}

func newReport(in Scrape) *Report {
	for _, repo := range in.Repos {
		fmt.Println(repo)
	}
	return &Report{
		in:  in.Repos,
		out: "",
	}
}
