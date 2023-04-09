package controller

import (
	"backend/cmd/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, _ := model.SqlStart()

	allUsers, err := model.GetRows(db)
	if err != nil {
		log.Fatalf("getRows error err*%v", err)
	}

	out, err := json.Marshal(allUsers)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	defer db.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
