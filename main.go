package main

import (
	"github.com/kijimaD/act/config"
	"github.com/kijimaD/act/gh"
	"github.com/kijimaD/act/report"
	"github.com/kijimaD/act/scrape"
)

func main() {
	config := config.NewConfig()
	config.Load()

	scrape := scrape.NewScrape(config).Execute()
	report.NewReport(*scrape, config).Execute()
	gh.New(config).Commit().Push()
}
