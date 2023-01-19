package hutils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func TestComparePassword(t *testing.T) {
	original := "TQeEuhvrEfbeIlYqBJmLjJWYVAbkVfJOMClDpWUcNtDtWjWRxI"
	hashed, err := bcrypt.GenerateFromPassword([]byte(original), bcrypt.DefaultCost)
	old := "$2a$10$iQHkXrTRBFSyg0LL6O8WbOa5PSJWgeZiuAQVr3hj2Tri3ZURwJEyO"
	if err != nil {
		t.Error(err)
	}

	log.Println(original)
	log.Println(string(hashed))
	log.Println(old)

	err = bcrypt.CompareHashAndPassword(hashed, []byte(original))
	if err != nil {
		t.Error(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(old), []byte(original))
	if err != nil {
		t.Error(err)
	}

	if !ComparePassword(original, string(old)) {
		t.Error("password not equal")
	}
}
