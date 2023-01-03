package entities

import (
	"github.com/Yegor-own/Chat/src/v1/internal/database"
	"testing"
)

func TestMigrate(t *testing.T) {
	db, err := database.ConnectPostgres()
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&User{}, &Relationship{}, &Message{}, &Chat{})
	if err != nil {
		t.Error(err)
	}
}
