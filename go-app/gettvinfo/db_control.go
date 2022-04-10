package gettvinfo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

// DBと接続する
func init() {
	// envファイルからpasswordとユーザー名を取得し、data source nameを作成
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatalln(err)
	}

	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@(localhost:3306)/%s?interpolateParams=true&parseTime=true", user, pw, dbName)

	// mysqlと接続
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("OpenError: ", err)
	}

	// 接続しているか確認
	if err := Db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
}

// 要素を挿入
func programInsert(program string) {
	// insert, 自動採番されるためidは不要
	const insertProgram = "INSERT INTO programs(program_name, created_at, updated_at) VALUES(?,?,?)"
	time := time.Now().Local()
	_, err := Db.Exec(insertProgram, program, time, time)
	if err != nil {
		log.Fatal("InsertError: ", err)
	}
}

// programがDBに登録されているか確認
func isProgram(program string) bool {
	row := Db.QueryRow(`select exists (select * from programs where program_name = ?)`, program)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		log.Println("Failed to scan exists from programs in isProgram ", err)
	}
	return exists
}

// episodeを挿入(番組名がある前提)
func episodeInsert(program string) {
	// 番組のIDを取得
	row := Db.QueryRow(`select id from programs where program_name = ?`, program)
	var id int
	if err := row.Scan(&id); err != nil {
		log.Print("rowScanError: ", err)
		return
	}

	// すでに同じ日付でepisodeレコードが作成されていないか検索
	time := time.Now().Local()
	day := time.Format("2006-01-02")
	row = Db.QueryRow(`select exists (select * from episodes where date = ? and program_id = ?)`, day, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		log.Fatal("rowScanError: ", err)
	}
	if exists {
		return
	}

	// episodeをinsertする
	const insertProgram = "INSERT INTO episodes(program_id, date, episode_number, episode_title, created_at, updated_at) VALUES(?,?,?,?,?,?)"
	_, err := Db.Exec(insertProgram, id, day, nil, nil, time, time)
	if err != nil {
		log.Fatal("InsertError: ", err)
	}
}