package main

import (
	"log"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func main() {

	const scrapeURL = "https://tver.jp/program"               // スクレイピングするurl
	const selector = ".epg-item_seriesTitleText__RnbO0" // htmlから抜き出す要素

	// スクレイピング
	programs := gettvinfo.Scraping(scrapeURL, selector)

	// 番組名と放送回をinsert
	if err := gettvinfo.ProgramInsert(programs); err != nil {
		log.Printf("Fatal: ProgramInsert %v", err)
	}

	if err := gettvinfo.EpisodeInsert(programs); err != nil {
		log.Printf("Fatal: EpisodeInsert %v", err)
	}
}
