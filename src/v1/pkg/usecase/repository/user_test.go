package repository

import (
	"github.com/Yegor-own/Chat/src/v1/internal/database"
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestNilStringCompare(t *testing.T) {
	var nilStr string = ""
	str := "not nil"
	log.Println(&nilStr == nil, &str == nil, reflect.ValueOf(nilStr).IsZero())
}

func TestCreateUser(t *testing.T) {
	db, err := database.ConnectPostgres()
	if err != nil {
		t.Error(err)
	}

	repo := NewUserRepository(db)
	user, err := repo.CreateUser("nn", "secret")
	if err != nil {
		t.Error(err)
	}
	log.Println(user)
}

func TestReadUser(t *testing.T) {
	db, err := database.ConnectPostgres()
	if err != nil {
		t.Error(err)
	}

	repo := NewUserRepository(db)
	user, err := repo.ReadUser(1)
	if err != nil {
		t.Error(err)
	}
	log.Println(user)
}

func TestUpdateUser(t *testing.T) {
	db, err := database.ConnectPostgres()
	if err != nil {
		t.Error(err)
	}

	var u entities.User
	u.ID = 5
	db.First(&u)

	ls := time.Now()
	u.LastSeen = &ls

	repo := NewUserRepository(db)
	user, err := repo.UpdateUser(&u)
	if err != nil {
		t.Error(err)
	}
	log.Println(user, u)
}
