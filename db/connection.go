package db

import (
	"bipbop/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbConfig.Host, config.DbConfig.Port, config.DbConfig.User, config.DbConfig.Password, config.DbConfig.Name )
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}