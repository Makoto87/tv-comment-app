package main

import (
	"log"

	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func main() {

	const url = "https://tver.jp/program"                   // スクレイピングするurl
	const selector = "span.epg-item_seriesTitleText__RnbO0" // htmlから抜き出す要素

	// htmlを取得する
	html, err := gettvinfo.Scraping(url, selector)
	if err != nil {
		log.Println("Failed to get html", err)
		return
	}

	// htmlから要素を抜き出し、DBへ保存
	if err = gettvinfo.SaveElements(html, selector); err != nil {
		log.Println(err)
	}
}
