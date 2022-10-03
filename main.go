package main

import (
	"act/config"
	"act/report"
	"act/scrape"
	"act/gh"
)

func main() {
	config := config.NewConfig()
	config.Load()

	scrape := scrape.NewScrape(config).Execute()
	report.NewReport(*scrape, config).Execute()
	gh.New(config).Commit().Push()
}
