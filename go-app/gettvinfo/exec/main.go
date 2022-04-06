package main

import (
	"github.com/Makoto87/tv-comment-app/go-app/gettvinfo"
)

func main() {

	const url = "https://tver.jp/program"                   // スクレイピングするurl
	const selector = "span.epg-item_seriesTitleText__RnbO0" // htmlから抜き出す要素

	// htmlを取得する
	var html string = gettvinfo.Scraping(url, selector)

	// htmlから要素を抜き出し、DBへ保存
	gettvinfo.SaveElements(html, selector)
}
