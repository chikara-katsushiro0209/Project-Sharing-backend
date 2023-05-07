package model

import (
	"backend/cmd/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Birthday  int
	Email     string
	Password  string
	Created   time.Time
	Updated   time.Time
}

// 全件取得
func GetRows(db *sql.DB) ([]User, error) {
	// db, err := SqlStart()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.Birthday, &u.Created, &u.Updated); err != nil {
			log.Fatalf("getRows rows.Scan error err*%v", err)
			return nil, err
		}

		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err*%v", err)
		return nil, err
	}
	// defer db.Close()
	return users, nil

}

// 条件取得
func GetSingleRow(db *sql.DB, userID int) (User, error) {
	user := &User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Birthday, &user.Created, &user.Updated)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("getSingleRow no records.")
	}

	if err != nil {
		log.Fatalf("getSingleRow db.QueryRow error err:%v", err)
	}

	return *user, nil
}

// 登録処理
func InsertUser(db *sql.DB, firstName, lastName string, birthday int, email string, password string) error {
	hashPassword, _ := domain.HashPassword(password)

	err := performTransaction(db, func(tx *sql.Tx) error {
		res, err := db.Exec(
			"INSERT INTO users (first_name,last_name,birthday,email,password) VALUES (?,?,?,?,?)",
			firstName,
			lastName,
			birthday,
			email,
			hashPassword,
		)

		if err != nil {
			log.Fatalf("insertUser db.Exec error %v", err)
		}

		_, err = res.LastInsertId()
		if err != nil {
			log.Fatalf("insetUser res.LastInsertId error %v", err)
		}

		return nil
	})

	return err
}

func UpdateUser(db *sql.DB, id, firstName, lastName string, password string) error {
	hashPassword, _ := domain.HashPassword(password)
	_, err := db.Exec("UPDATE users SET first_Name = ?, last_Name = ?, password = ? WHERE id = ?", firstName, lastName, hashPassword, id)
	if err != nil {
		log.Fatalf("updateUser db.Exec error %v", err)
		return err
	}
	return nil
}
