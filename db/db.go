package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rodiond26/gomigrator/config"
)

// NewDB .
func NewDB(cfg *config.Config) *sql.DB {
	fmt.Println("Connecting to Postgres database...")
	fmt.Printf(">>> cfg [%+v]\n", cfg)
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.Db["user"], cfg.Db["password"], cfg.Host, cfg.Port, cfg.Db["name"])
	fmt.Println(">>> dsn = ", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Unable to connect to database 1", err.Error())
		return nil
	}
	// defer db.Close()

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
