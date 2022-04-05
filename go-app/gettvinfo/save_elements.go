package gettvinfo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SaveElements(html string, selector string) {

	// envファイルからpasswordとユーザー名を取得し、data source nameを作成
	if err := godotenv.Load("go.env"); err != nil {
		log.Fatalln(err)
	}

	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@(localhost:3306)/%s?charset=utf8", user, pw, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}

	return

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
