package gettvinfo

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

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

// 番組を挿入
func ProgramInsert(programs []string) error {
	insert := "INSERT IGNORE INTO programs(program_name, created_at, updated_at) VALUES "

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

// episodeを挿入
func EpisodeInsert(programs []string) error {
	// sliceを文字列に変換
	var programStr string
	for _, p := range programs {
		programStr += fmt.Sprintf(`"%s",`, p)
	}
	programStr = programStr[:len(programStr)-1]

	// 番組のIDをまとめて取得
	rows, err := Db.Query(`select id from programs where program_name in (` + programStr + `)`)
	if err != nil {
		return fmt.Errorf(" failed to Query id from programs where program_name in args: %w", err)
	}
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return fmt.Errorf(" failed to rows.Scan id from programs: %w", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf(" rows.Err : %w", err)
	}

	// episodeをinsertする
	insert := "INSERT IGNORE INTO episodes(program_id, date, episode_number, episode_title, created_at, updated_at) VALUES "

	for _, id := range ids {
		value := fmt.Sprintf(`(%v, NOW(), null, null, NOW(), NOW()),`, id)
		insert += value
	}

	insert = insert[:len(insert)-1]

	if _, err := Db.Exec(insert); err != nil {
		return fmt.Errorf(" failed to insert episode: %w", err)
	}
	return nil
}
