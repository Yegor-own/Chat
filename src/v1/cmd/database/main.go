package main

import (
	"github.com/Yegor-own/Chat/src/v1/internal/database"
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/Yegor-own/Chat/src/v1/pkg/factories"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

	log.Println("Я живой")
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connection success")

	err = db.AutoMigrate(&entities.User{}, &entities.Chat{}, &entities.Message{}, &entities.Relationship{})
	if err != nil {
		log.Fatalln(err)
	}

	users, err := factories.UserFactory(5)
	if err != nil {
		log.Println("failed to generate users")
	}

	res := db.Create(&users)
	if res.Error != nil || res.RowsAffected <= 0 {
		log.Println("failed to create users")
	}

	log.Println(res.RowsAffected)

	u := []entities.User{}
	res = db.Find(&u)
	if res.Error != nil || res.RowsAffected <= 0 {
		log.Println("failed to find users")
	}

	log.Println(res.RowsAffected)
}
