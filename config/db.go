package config

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func ConnectDb() *sql.DB {
	
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	
	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
	)
	
	var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}