package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func GetDatabaseConnection() *sql.DB {

	var err error
	if db != nil {
		return db
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(time.Second * 60)
	db.SetConnMaxLifetime(time.Minute * 5)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to open database: %v", err))
	} else if err := db.Ping(); err != nil {
		slog.Error(fmt.Sprintf("failed to connect to database: %v", err))
	}

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	return db
}
