package main

import ()

func main() {
	config := newConfig()
	config.load(DefaultConfigFilePath)

	scrape := newScrape(config)
	newReport(scrape, config).execute()
}
