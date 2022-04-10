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
var nowFunc func() time.Time

func init() {
	setDb()  // DBの設定
	setJST() // JST取得のための設定
}

// DBの初期設定
func setDb() {
	// envファイルがあれば読み込む
	godotenv.Load("go.env")

	// 環境変数からMySqlのuser名・password・DB名を取得し、dataSourceNameを作成
	user, okUser := os.LookupEnv("MYSQL_USER")
	pw, okPw := os.LookupEnv("MYSQL_PASSWORD")
	dn, okDn := os.LookupEnv("DATABASE_NAME")
	if !okUser || !okPw || !okDn {
		log.Fatalf("Fataled to get value for DB. user: %v, password: %v, DB name: %v", okUser, okPw, okDn)
	}
	dsn := fmt.Sprintf("%s:%s@(localhost:3306)/%s?interpolateParams=true&parseTime=true", user, pw, dn)

	// mysqlと接続
	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("OpenError: ", err)
	}

	// 接続しているか確認
	if err := Db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
}

// JSTを取得できる関数を設定
func setJST() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("can't get Asia/Tokyo: ", err)
	}

	nowFunc = func() time.Time {
		return time.Now().In(jst)
	}
}

// 要素を挿入
func programInsert(program string) error {
	// insert, 自動採番されるためidは不要
	const insertProgram = "INSERT INTO programs(program_name, created_at, updated_at) VALUES(?,?,?)"
	time := nowFunc()
	if _, err := Db.Exec(insertProgram, program, time, time); err != nil {
		return fmt.Errorf(" failed to insert program (%s): %w", program, err)
	}
	return nil
}

// programがDBに登録されているか確認
func isProgram(program string) (bool, error) {
	row := Db.QueryRow(`select exists (select * from programs where program_name = ?)`, program)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf(" failed to scan %s exists from programs: %w", program, err)
	}
	return exists, nil
}

// episodeを挿入(番組名がある前提)
func episodeInsert(program string) error {
	// 番組のIDを取得
	row := Db.QueryRow(`select id from programs where program_name = ?`, program)
	var id int
	if err := row.Scan(&id); err != nil {
		log.Print("rowScanError: ", err)
		return fmt.Errorf(" failed to scan %s id from programs: %w", program, err)
	}

	// すでに同じ日付でepisodeレコードが作成されていないか検索
	time := nowFunc()
	day := time.Format("2006-01-02")
	row = Db.QueryRow(`select exists (select * from episodes where date = ? and program_id = ?)`, day, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return fmt.Errorf(" failed to scan * from episodes where date and program_id in %s, %v: %w", program, day, err)
	}
	if exists {
		fmt.Printf("Skip to insert episode because episode is exist in %s, %v \n", program, day)
		return nil
	}

	// episodeをinsertする
	const insertProgram = "INSERT INTO episodes(program_id, date, episode_number, episode_title, created_at, updated_at) VALUES(?,?,?,?,?,?)"
	_, err := Db.Exec(insertProgram, id, day, nil, nil, time, time)
	if err != nil {
		return fmt.Errorf(" failed to insert episode (%s, %v): %w", program, day, err)
	}

	return nil
}
