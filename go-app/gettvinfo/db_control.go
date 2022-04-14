package gettvinfo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "time/tzdata"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB
var nowFunc func() time.Time

func init() {
	setDb() // DBの設定
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
	dsn += "&loc=Asia%2FTokyo"

	// mysqlと接続
	var err error
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
func ProgramInsert(programs []string) error {
	insert := "INSERT IGNORE INTO tests(program_name, created_at, updated_at) VALUES "
	// insert := "INSERT IGNORE INTO programs(program_name, created_at, updated_at) VALUES "

	for _, p := range programs {
		value := fmt.Sprintf(`("%s", NOW(), NOW()),`, p)
		insert += value
	}
	insert = insert[:len(insert)-1]

	if _, err := Db.Exec(insert); err != nil {
		return fmt.Errorf(" failed to insert program: %w", err)
	}
	return nil
}

// episodeを挿入(番組名がある前提)
func episodeInsert(program string) error {
	// 番組のIDを取得
	row := Db.QueryRow(`select id from programs where program_name = ?`, program)
	var id int
	if err := row.Scan(&id); err != nil {
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
	fmt.Println("episode insert success ", program)

	return nil
}
