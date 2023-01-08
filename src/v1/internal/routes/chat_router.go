package routes

import (
	"github.com/Yegor-own/Chat/src/v1/internal/controllers"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
)

func ChatRouter(app fiber.Router, service service.ChatService, relationship service.RelationshipService) {
	app.Get("/", controllers.GetChat(service))
	app.Post("/", controllers.CreateChat(service))
	app.Put("/", controllers.UpdateChat(service))
	app.Delete("/", controllers.DeleteChat(service))
}
