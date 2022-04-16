package main

import (
	"log"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func main() {

	const url = "https://tver.jp/program"               // スクレイピングするurl
	const selector = ".epg-item_seriesTitleText__RnbO0" // htmlから抜き出す要素

	// スクレイピング
	programs := gettvinfo.Scraping(url, selector)

	// 番組名と放送回をinsert
	if err := gettvinfo.ProgramInsert(programs); err != nil {
		log.Println("Fatal: ProgramInsert ", err)
	}

	if err := gettvinfo.EpisodeInsert(programs); err != nil {
		log.Println("Fatal: EpisodeInsert ", err)
	}
}
