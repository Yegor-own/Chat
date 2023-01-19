package routes

import (
	"github.com/Yegor-own/Chat/src/v1/internal/controllers"
	"github.com/Yegor-own/Chat/src/v1/internal/middleware"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
)

func ChatRouter(app fiber.Router, service service.ChatService, relationship service.RelationshipService) {
	app.Get("/", middleware.Protected(), controllers.GetChat(service))
	app.Post("/", middleware.Protected(), controllers.CreateChat(service))
	app.Put("/", middleware.Protected(), controllers.UpdateChat(service))
	app.Delete("/", middleware.Protected(), controllers.DeleteChat(service))
}
