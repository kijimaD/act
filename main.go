package main

import (
	"act/config"
	"act/report"
	"act/scrape"
)

func main() {
	config := config.NewConfig()
	config.Load()

	scrape := scrape.NewScrape(config)
	report.NewReport(scrape, config).Execute()
}
