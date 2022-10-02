package main

import ()

func main() {
	config := newConfig()
	config.load(DefaultConfigFilePath)

	scrape := newScrape()
	newReport(scrape).execute(config)
}
