package domain

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatalf("HashPassword bcrypt.GenerateFromPassword err:%v", err)
		return nil, err
	}

	return hashedPassword, nil
}

func ConfirmPassword(password, hashedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		log.Printf("ConfirmPassword bcrypt.ConfirmPassword err:%v", err)
		return false
	}

	return true
}
