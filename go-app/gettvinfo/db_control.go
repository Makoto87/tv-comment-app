// get tv program information by scraping, and insert information into database
package gettvinfo

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	setDB()
}

// set up to connect database
func setDB() {
	godotenv.Load("go.env")

	user, okUser := os.LookupEnv("MYSQL_USER")
	pw, okPW := os.LookupEnv("MYSQL_PASSWORD")
	host, okHost := os.LookupEnv("HOST")
	dn, okDN := os.LookupEnv("DATABASE_NAME")
	if !okUser || !okPW || !okDN {
		log.Fatalf("Fataled to get value for DB. user: %v, password: %v, host: %v, DB name: %v", okUser, okPW, okHost, okDN)
	}
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/%s?interpolateParams=true&parseTime=true&loc=Asia%%2FTokyo", user, pw, host, dn)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("OpenError: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
}

// insert program into database. any number is acceptable.
func ProgramInsert(programs []string) error {
	insert := "INSERT IGNORE INTO programs(program_name, created_at, updated_at) VALUES "

	vals := make([]any, 0, len(programs))
	for _, p := range programs {
		insert += `(?, NOW(), NOW()),`
		vals = append(vals, p)
	}
	insert = insert[:len(insert)-1]

	stmt, err := DB.Prepare(insert)
	if err != nil {
		return fmt.Errorf("failed to prepare insert program: %w", err)
	}

	if _, err := stmt.Exec(vals...); err != nil {
		return fmt.Errorf("failed to insert program: %w", err)
	}
	return nil
}

// insert episodes into database. First, get all id from programs. Second, insert episodes which use id from programs.
func EpisodeInsert(programs []string) error {
	selectID := `select id from programs where program_name in (`
	vals := make([]any, 0, len(programs))
	for _, p := range programs {
		selectID += `?,`
		vals = append(vals, p)
	}
	selectID = selectID[:len(selectID)-1] + `)`

	stmt, err := DB.Prepare(selectID)
	if err != nil {
		return fmt.Errorf("failed to prepare select program id: %w", err)
	}

	rows, err := stmt.Query(vals...)
	if err != nil {
		return fmt.Errorf("failed to Query id from programs where program_name in args: %w", err)
	}
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return fmt.Errorf("failed rows.Scan id from programs: %w", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows.Err : %w", err)
	}

	insert := "INSERT IGNORE INTO episodes(program_id, date, episode_number, episode_title, created_at, updated_at) VALUES "

	vals = make([]any, 0, len(ids))
	for _, id := range ids {
		insert += `(?, NOW(), null, null, NOW(), NOW()),`
		vals = append(vals, id)
	}
	insert = insert[:len(insert)-1]

	stmt, err = DB.Prepare(insert)
	if err != nil {
		return fmt.Errorf("failed to prepare insert episode: %w", err)
	}

	if _, err := stmt.Exec(vals...); err != nil {
		return fmt.Errorf("failed to insert episode: %w", err)
	}
	return nil
}
