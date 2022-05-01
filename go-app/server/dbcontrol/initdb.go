package dbcontrol

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// set up db connection
func init() {
	godotenv.Load("go.env")

	user, okUser := os.LookupEnv("MYSQL_USER")
	pw, okPW := os.LookupEnv("MYSQL_PASSWORD")
	host, okHost := os.LookupEnv("HOST")
	dn, okDN := os.LookupEnv("DATABASE_NAME")
	if !okUser || !okPW || !okHost || !okDN {
		log.Fatalf("[server/graph/dbcontrol] Fataled to get value for DB. user: %v, password: %v, host: %v, DB name: %v", okUser, okPW, okHost, okDN)
	}
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/%s?interpolateParams=true&parseTime=true&loc=Asia%%2FTokyo", user, pw, host, dn)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("[server/graph/dbcontrol] OpenError: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("[server/graph/dbcontrol] PingError: ", err)
	}
}
