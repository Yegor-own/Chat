package main

import (
	"github.com/Yegor-own/Chat/src/v1/internal/database"
	"github.com/Yegor-own/Chat/src/v1/internal/routes"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/repository"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connection success")
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api/v1")
	routes.UserRouter(api, userService)
	log.Fatal(app.Listen(":8080"))
}
