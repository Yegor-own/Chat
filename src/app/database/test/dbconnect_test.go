package test

import (
	"chat/app/database"
	"testing"
)

func TestDBConnect(t *testing.T) {
	params, err := database.LoadEnv()
	if err != nil {
		t.Error(err)
	}

	db, err := database.ConnectDB(*params)
	if err != nil {
		t.Error(err.Error())
	}
	if db == nil {
		t.Error("db connection empty")
	}
}
