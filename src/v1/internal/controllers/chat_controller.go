package controllers

import (
	"github.com/Yegor-own/Chat/src/v1/internal/interfaces"
	"github.com/Yegor-own/Chat/src/v1/internal/presenters"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetChat(service service.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var chatId interfaces.IdChat
		err := ctx.BodyParser(&chatId)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		res, err := service.GetChat(chatId.Id)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func CreateChat(service service.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var chatData interfaces.CreateChat
		err := ctx.BodyParser(&chatData)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		res, err := service.InsertChat(chatData.Title)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func UpdateChat(service service.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var chatData interfaces.UpdateChat
		err := ctx.BodyParser(&chatData)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		res, err := service.UpdateChat(chatData.Id, chatData.Title)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest(res))
	}
}

func DeleteChat(service service.ChatService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var chatId interfaces.IdChat
		err := ctx.BodyParser(&chatId)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		err = service.RemoveChat(chatId.Id)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(presenters.BadResponse(err))
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(presenters.SuccessRequest("Deleted chat"))
	}
}
