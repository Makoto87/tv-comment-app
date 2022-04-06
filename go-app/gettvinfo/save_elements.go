package gettvinfo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Program struct {
	id                   int
	programName          string
	createdAt, updatedAt time.Time
}

func SaveElements(html string, selector string) {

	// envファイルからpasswordとユーザー名を取得し、data source nameを作成
	if err := godotenv.Load("go.env"); err != nil {
		log.Fatalln(err)
	}

	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@(localhost:3306)/%s?interpolateParams=true&parseTime=true", user, pw, dbName)

	// mysqlと接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	defer db.Close()

	// 接続しているか確認
	if err := db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}

	// 同じ番組名があればINSERTはスキップ
	row := db.QueryRow(`select exists (select * from programs where program_name = ?)  `, "あちこち")
	var exists bool
	if err = row.Scan(&exists); err != nil {
		log.Fatal("rowScanError: ", err)
	}
	if exists == true { return }

	// insert, 自動採番されるためidは不要
	const insertProgram = "INSERT INTO programs(program_name, created_at, updated_at) VALUES(?,?,?)"
	time := time.Now()
	result, err := db.Exec(insertProgram, "test2", time, time)
	if err != nil {
		log.Fatal("InsertError: ", err)
	}
	
	// id取得
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("LastInsertIdError: ", err)
	}
	fmt.Println(id)

	// SELECT
	rows, err := db.Query(`select * from programs order by id asc`)
	if err != nil {
		log.Fatal("QueryError: ", err)
	}
	for rows.Next() {
		var p Program
		if err := rows.Scan(&p.id, &p.programName, &p.createdAt, &p.updatedAt); err != nil {
			log.Fatalln("rowsScan: ", err)
		}
		fmt.Println(p)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln("rowsErr: ", err)
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
