package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type Database struct {
	DB *sql.DB
}

func New() (Database, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Database{}, err
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return Database{}, err
	}

	time.Sleep(time.Second * 5)
	if err = db.Ping(); err != nil {
		return Database{}, err
	}

	return Database{DB: db}, nil
}
