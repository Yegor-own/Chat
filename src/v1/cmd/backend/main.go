package main

import (
	"github.com/Yegor-own/Chat/src/v1/internal/database"
	"github.com/Yegor-own/Chat/src/v1/internal/routes"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/repository"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	chatRepo := repository.NewChatRepository(db)
	chatService := service.NewChatService(chatRepo)

	relationshipRepo := repository.NewRelationshipRepositorySQL(db)
	relationshipService := service.NewRelationshipService(relationshipRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api/v1")
	userGroup := api.Group("/user")
	chatGroup := api.Group("/chat")

	routes.UserRouter(userGroup, userService, relationshipService)
	routes.ChatRouter(chatGroup, chatService, relationshipService)

	log.Fatal(app.Listen(":8080"))
}
