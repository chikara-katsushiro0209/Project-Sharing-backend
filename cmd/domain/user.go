package domain

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("HashPassword bcrypt.GenerateFromPassword err:%v", err)
	}

	return string(hashedPassword), nil
}

func ConfirmPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Fatalf("HashPassword bcrypt.ConfirmPassword err:%v", err)
		return false
	}

	return true
}
