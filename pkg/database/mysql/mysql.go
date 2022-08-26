package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// siakad
var DB2 *sqlx.DB

// DBSIA open connection database
func DBSIA() {

	var (
		host = os.Getenv("SIA_DB_HOST")
		port = os.Getenv("SIA_DB_PORT")
		user = os.Getenv("SIA_DB_USER")
		pass = os.Getenv("SIA_DB_PASS")
		name = os.Getenv("SIA_DB_NAME")
	)

	mysqlInfo := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, name)

	db, err := sqlx.Open("mysql", mysqlInfo)
	if err = db.Ping(); err != nil {
		log.Println("[SIA] An error occured while sending PING to database! ", err)
	} else {
		log.Println("Connected to database 'sia'!")
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DB2 = db
}
