package getTVInformation

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func Scraping() {
	// fmt.Println("start scraping")
	// doc, err := goquery.NewDocument("https://tv.yahoo.co.jp/search?q=tomato")
	// // doc, err := goquery.NewDocument("https://rooter.jp/web-crawling/go_goquery_scraping/")

	// if err != nil {
	// 	fmt.Println("document not found. ")
	// 	os.Exit(1)
	// }

	// fmt.Println("After goquery document")

	// program := ""
	// doc.Find(".globalNavigationText").Each(func(_ int, s *goquery.Selection) {
	// // doc.Find(".prettyprint").Each(func(_ int, s *goquery.Selection) {
	// 	fmt.Println("find now")
	// 	program += s.Text() + "\n"
	// })

	// fmt.Println("after doc.Find")

	// fmt.Println(program)

	// ブラウザはChromeを指定して起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	// ログインページに遷移
	if err := page.Navigate("https://qiita.com/login"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	// ID, Passの要素を取得し、値を設定
	identity := page.FindByID("identity")
	password := page.FindByID("password")
	identity.Fill("Your Id Here.")
	password.Fill("Your Passowrd Here.")
	// formをサブミット
	if err := page.FindByClass("loginSessionsForm_submit").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}
	// 処理完了後、3秒間ブラウザを表示しておく
	time.Sleep(3 * time.Second)
}
