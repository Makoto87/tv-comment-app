package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// set up db connection
func init() {
	godotenv.Load("go.env")

	user, okUser := os.LookupEnv("MYSQL_USER")
	pw, okPW := os.LookupEnv("MYSQL_PASSWORD")
	host, okHost := os.LookupEnv("HOST")
	dn, okDN := os.LookupEnv("DATABASE_NAME")
	if !okUser || !okPW || !okHost || !okDN {
		log.Fatalf("[server/graph/model] Fataled to get value for DB. user: %v, password: %v, host: %v, DB name: %v", okUser, okPW, okHost, okDN)
	}
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/%s?interpolateParams=true&parseTime=true&loc=Asia%%2FTokyo", user, pw, host, dn)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("[server/graph/model] OpenError: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("[server/graph/model] PingError: ", err)
	}
}

func GetPrograms(search string) ([]*Program, error) {
	var programs []*Program

	query := "select id, program_name from programs where program_name like ?"
	if search == "" {
		search = "%"
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select id & program_name: %w", err)
	}

	rows, err := stmt.Query(search)
	if err != nil {
		return nil, fmt.Errorf("failed to select program id, program_name by Query: %w", err)
	}

	for rows.Next() {
		var p Program
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, fmt.Errorf("failed rows.Scan id & program_name from program: %w", err)
		}
		programs = append(programs, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Error: %w", err)
	}

	return programs, nil
}
