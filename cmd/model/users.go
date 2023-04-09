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
func GetSingleRow(db *sql.DB, userID int) {
	u := &User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Birthday, &u.Email, &u.Password, &u.Created, &u.Updated)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("getSingleRow no records.")
		return
	}

	if err != nil {
		log.Fatalf("getSingleRow db.QueryRow error err:%v", err)
	}
	fmt.Println(u)
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
