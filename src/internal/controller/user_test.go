package controller

import (
	"log"
	"testing"
)

func TestUserController(t *testing.T) {
	var user UserController

	res, err := user.NewUser("Xui", "nn")
	if err != nil {
		t.Error(err)
	}

	log.Println(res)
}
