package config

import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "wahyujatirestu"
	password = "password"
	dbname = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func ConnectDb() *sql.DB {
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