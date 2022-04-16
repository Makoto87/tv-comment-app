package gettvinfo

import (
	"time"

	"github.com/go-rod/rod"
)

func Scraping(url, selector string) []string {

	// rodを使ってヘッドレスブラウザを立ち上げる
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// URLを開く
	page := browser.MustPage(url)

	// 画面操作
	controllView(page)

	// 要素を取得
	var programs []string
	sections := page.MustElements(selector)
	for _, section := range sections {
		s := section.MustText()
		if s != "" {
			programs = append(programs, s)
		}
	}

	return programs
}

func controllView(page *rod.Page) {
	// 20秒間URLを読み込めなかったら処理は強制終了
	// MustWaitLoadでページが完全に表示されるまで待つ
	page.Timeout(20 * time.Second).MustWaitLoad().CancelTimeout()

	// 画面にあるボタンをクリックする
	page.MustElementR("button", "同意する").MustClick()
	page.MustElementR("button", "スキップ").MustWaitStable().MustClick()
	page.MustElement("button.button_button__GOl5m.modal_closeButton__4N3wA").MustWaitStable().MustClick()
}