package controller

import (
	"backend/cmd/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
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
	// vars := mux.Vars(r)
	// userId := vars["id"]
	userId := chi.URLParam(r, "id")
	fmt.Println("userId:", userId)

	var data struct {
		ID       string `json:"id"`
		LastName string `json:"lastName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Println("userId:", data.ID)
	data.ID = userId // Set the ID field to the value of userId
	fmt.Println("userId:", data.ID)
	fmt.Println("lastName:", data.LastName)

	out, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
