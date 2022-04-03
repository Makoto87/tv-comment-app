package getTVInformation

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func Scraping() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://example.com"),
		chromedp.Sleep(time.Second),
		chromedp.Text(`h1`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("h1 contains: '%s'\n", res)
}
