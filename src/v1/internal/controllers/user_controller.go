package controllers

import (
	"github.com/Yegor-own/Chat/src/v1/internal/interfaces"
	"github.com/Yegor-own/Chat/src/v1/internal/presenters"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetUser(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var userId interfaces.GetUser
		err := ctx.BodyParser(&userId)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.UserBadResponse(err))
		}
		res, err := service.GetUser(userId.Id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.UserBadResponse(err))
		}
		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.UserSuccessRequest(res))
	}
}
