package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "host=127.0.0.1 port=10000 user=postgres password=00000000 dbname=trabea sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return conn, nil
}
