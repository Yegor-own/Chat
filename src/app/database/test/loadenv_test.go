package test

import (
	"chat/app/database"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	params, err := database.LoadEnv()
	if err != nil {
		t.Error(err)
	}
	if params == nil {
		t.Error("params is empty")
	}
}
