package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// NewDB .
func NewDB() *sql.DB { // TODO fix
	fmt.Println("Connecting to Postgres database...")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println("psqlconn", psqlconn)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Unable to connect to database 1", err.Error())
		return nil
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Unable to connect to database 2", err.Error())
		return nil
	}

	fmt.Println("Database connected!")

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
