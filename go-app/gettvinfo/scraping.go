package gettvinfo

import (
	"time"

	"github.com/go-rod/rod"
)

// get all texts of selectors from url, return texts as string slice
func Scraping(url, selector string) []string {

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url)
	controllView(page)

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
	page.Timeout(20 * time.Second).MustWaitLoad().CancelTimeout()

	page.MustElementR("button", "同意する").MustClick()
	page.MustElementR("button", "スキップ").MustWaitStable().MustClick()
	page.MustElement("button.button_button__GOl5m.modal_closeButton__4N3wA").MustWaitStable().MustClick()
}