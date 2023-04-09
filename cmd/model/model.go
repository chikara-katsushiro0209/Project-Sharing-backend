package model

import (
	"database/sql"
	"log"
)

var db *sql.DB

func SqlStart() (*sql.DB, error) {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:13306)/db-go-db?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
		return nil, err
	}

	return db, nil
}
