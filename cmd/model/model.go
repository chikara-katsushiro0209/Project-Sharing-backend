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

// トランザクション
func performTransaction(db *sql.DB, txFunc func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return err
}
