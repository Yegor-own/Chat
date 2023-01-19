package factories

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/Yegor-own/Chat/src/v1/pkg/hutils"
	"github.com/bxcodec/faker/v4"
	"log"
)

func UserFactory(num int) ([]entities.User, error) {
	var users []entities.User

	for i := 0; i < num; i++ {
		name, email, password := faker.Name(), faker.Email(), faker.Password()
		log.Println("new user", name, email, password)

		hashed, err := hutils.HashPassword(password)
		if err != nil {
			return nil, err
		}

		users = append(users, entities.User{
			Name:     name,
			Email:    email,
			Password: hashed,
		})
	}

	return users, nil
}
