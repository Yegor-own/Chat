package routes

import (
	"github.com/Yegor-own/Chat/src/v1/internal/controllers"
	"github.com/Yegor-own/Chat/src/v1/internal/middleware"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service service.UserService, relationship service.RelationshipService) {
	app.Post("/login", controllers.Login(service))
	app.Get("/", controllers.GetUser(service))
	app.Post("/", controllers.CreateUser(service))
	app.Put("/", middleware.Protected(), controllers.UpdateUser(service))
	app.Delete("/", middleware.Protected(), controllers.DeleteUser(service))
}
