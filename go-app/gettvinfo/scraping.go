package gettvinfo

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func Scraping(url, selector string) (string, error) {

	// contextの用意
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// htmlの取得
	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(selector),
		// HTMLの取得
		chromedp.OuterHTML("html", &html, chromedp.ByQuery),
	); err != nil {
		return "", err
	}

	return html, nil
}
