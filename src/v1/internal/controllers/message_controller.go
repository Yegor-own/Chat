package controllers

import (
	"github.com/Yegor-own/Chat/src/v1/internal/interfaces"
	"github.com/Yegor-own/Chat/src/v1/internal/presenters"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateMessage(service service.MessageService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var msgData interfaces.CreateMessage
		err := ctx.BodyParser(&msgData)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		res, err := service.InsertMessage(msgData.Text, msgData.Id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func GetMessage(service service.MessageService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var msgData interfaces.GetMessage
		err := ctx.BodyParser(&msgData)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		res, err := service.GetMessage(msgData.Id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func GetAllMessages(service service.MessageService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var msgData interfaces.GetMessage
		err := ctx.BodyParser(&msgData)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		res, err := service.GetAllMessages(msgData.Id)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}
