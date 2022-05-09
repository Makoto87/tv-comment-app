package main

import (
	"log"
	"os"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("go.env")

	scrapeURL, okURL := os.LookupEnv("SCRAPE_URL")
	if !okURL {
		log.Fatalf("Fataled to get URL from environment file")
	}

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
