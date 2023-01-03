package presenters

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func UserBadResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func UserSuccessRequest(data *entities.User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}
