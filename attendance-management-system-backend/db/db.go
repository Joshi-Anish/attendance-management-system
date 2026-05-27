package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {

	var err error

	DB, err = sql.Open("mysql", "appuser:1234@tcp(127.0.0.1:3306)/attendance_db")
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	//  Connection pool settings (IMPORTANT)
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	//  Test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("MySQL connected successfully 🚀")
}
