package routes

import (
	"github.com/Yegor-own/Chat/src/v1/internal/controllers"
	"github.com/Yegor-own/Chat/src/v1/internal/middleware"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
)

func MessageRouter(app fiber.Router, service service.MessageService) {
	app.Post("/", middleware.Protected(), controllers.CreateMessage(service))
	app.Get("/", middleware.Protected(), controllers.GetMessage(service))
	app.Get("/all", middleware.Protected(), controllers.GetAllMessages(service))
}
