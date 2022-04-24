package main

import (
	"log"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func main() {

	const scrapeURL = "https://tver.jp/program"
	const selector = ".epg-item_seriesTitleText__RnbO0"

	programs := gettvinfo.Scraping(scrapeURL, selector)

	// TODO(makoto): function to get information of episode number and title

	if err := gettvinfo.ProgramInsert(programs); err != nil {
		log.Printf("Fatal: ProgramInsert: %v", err)
	}

	if err := gettvinfo.EpisodeInsert(programs); err != nil {
		log.Printf("Fatal: EpisodeInsert: %v", err)
	}
}
