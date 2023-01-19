package factories

import (
	"log"
	"testing"
)

func TestUserFactory(t *testing.T) {
	users, err := UserFactory(10)
	if err != nil {
		t.Error(err)
	}

	log.Println(users)
}
