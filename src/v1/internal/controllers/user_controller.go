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
		var userId interfaces.IdUser
		err := ctx.BodyParser(&userId)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		res, err := service.GetUser(userId.Id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func CreateUser(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var userData interfaces.CreateUser
		err := ctx.BodyParser(&userData)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}
		res, err := service.InsertUser(userData.Name, userData.Password)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.BadResponse(err))
		}
		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func UpdateUser(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var userData interfaces.UpdateUser
		err := ctx.BodyParser(&userData)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}
		u, err := service.GetUser(userData.ID)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}
		u.ID = userData.ID
		interfaces.MixModels(u, userData)

		res, err := service.UpdateUser(u)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.BadResponse(err))
		}
		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func DeleteUser(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var userId interfaces.IdUser
		err := ctx.BodyParser(&userId)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}
		err = service.RemoveUser(userId.Id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.BadResponse(err))
		}
		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest("Deleted user"))
	}
}
