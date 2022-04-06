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
func insert(program string) {
	// 同じ番組名があればINSERTはスキップ
	row := Db.QueryRow(`select exists (select * from programs where program_name = ?)  `, program)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		log.Fatal("rowScanError: ", err)
	}
	if exists == true {
		return
	}

	// insert, 自動採番されるためidは不要
	const insertProgram = "INSERT INTO programs(program_name, created_at, updated_at) VALUES(?,?,?)"
	time := time.Now()
	_, err := Db.Exec(insertProgram, program, time, time)
	if err != nil {
		log.Fatal("InsertError: ", err)
	}
}
