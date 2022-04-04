package gettvinfo

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SaveElements(html string, selector string) {
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
