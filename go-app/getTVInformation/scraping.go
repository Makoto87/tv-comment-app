package getTVInformation

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func Scraping() {

	// contextの用意
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// htmlの取得
	var selector = "span.epg-item_seriesTitleText__RnbO0"
	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://tver.jp/program"),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(selector),
		// HTMLの取得
		chromedp.OuterHTML("html", &html, chromedp.ByQuery),
	); err != nil {
		log.Fatalln(err)
	}

	// HTMLをDOMオブジェクトへ変換
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	// DOMノードを検索し、空欄以外全て取得
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str := selection.Text()
		if str != "" {
			fmt.Println(str)
		}
	})
}
