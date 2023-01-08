package main

import (
	"github.com/Yegor-own/Chat/src/v1/internal/database"
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"log"
)

func main() {
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
}
