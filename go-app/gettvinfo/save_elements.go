package gettvinfo

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SaveElements(html string, selector string) error {

	// HTMLをDOMオブジェクトへ変換
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return fmt.Errorf("failed to change HTML to DOM object: %w", err)
	}

	// DOMノードを検索し、空欄以外全て取得
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str := selection.Text()
		if str != "" {
			// program nameが登録されているか確認
			exists, err := isProgram(str)
			if err != nil {
				log.Println("Failed isProgram: ",err)
				return	// エラーの場合は次のselectorへ
			}

			// program nameがなかったらDBに登録
			if exists {
				if err = programInsert(str); err != nil {
					log.Println("Failed programInsert: ", err)
					return	// エラーの場合は次のselectorへ
				}
			}

			// episodeをDBに登録
			if err = episodeInsert(str); err != nil {
				log.Println("Failed episodeInsert: ", err)
			}
		}
	})

	return nil
}
