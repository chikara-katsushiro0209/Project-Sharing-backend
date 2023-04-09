package main

import (
	// "backend/cmd/model"

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
