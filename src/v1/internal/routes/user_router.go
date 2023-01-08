package routes

import (
	"github.com/Yegor-own/Chat/src/v1/internal/controllers"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service service.UserService, relationship service.RelationshipService) {
	app.Get("/", controllers.GetUser(service))
	app.Post("/", controllers.CreateUser(service))
	app.Put("/", controllers.UpdateUser(service))
	app.Delete("/", controllers.DeleteUser(service))
}
