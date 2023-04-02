package main

import (
	// "backend/cmd/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	var app application

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

func sqlStart() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:13306)/db-go-db?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}

	defer db.Close()

	// model.GetRows(db)
	// getSingleRow(db, 1)

	// insertUser(db, "鈴木", "太郎", 19900922)
}
