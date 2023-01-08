package routes

import (
	"github.com/Yegor-own/Chat/src/v1/internal/controllers"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
)

func MessageRouter(app fiber.Router, service service.MessageService) {
	app.Post("/", controllers.CreateMessage(service))
	app.Get("/", controllers.GetMessage(service))
	app.Get("/all", controllers.GetAllMessages(service))
}
