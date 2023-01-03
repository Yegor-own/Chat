package routes

import (
	"github.com/Yegor-own/Chat/src/v1/internal/controllers"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service service.UserService) {
	app.Get("/user", controllers.GetUser(service))
}
