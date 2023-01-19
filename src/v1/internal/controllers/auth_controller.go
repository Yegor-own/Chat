package controllers

import (
	"errors"
	"github.com/Yegor-own/Chat/src/v1/internal/interfaces"
	"github.com/Yegor-own/Chat/src/v1/internal/presenters"
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/Yegor-own/Chat/src/v1/pkg/hutils"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

func Login(service service.UserService) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		input := new(interfaces.LoginInput)
		userData := new(entities.User)

		err := ctx.BodyParser(input)
		if err != nil {
			return ctx.JSON(presenters.BadResponse(err))
		}

		userData, err = service.GetUserByEmail(input.Email)
		if err != nil {
			return ctx.JSON(presenters.BadResponse(err))
		}

		if userData.Name == "" {
			return ctx.JSON(presenters.BadResponse(errors.New("there is no user found")))
		}

		log.Println(userData)

		if !hutils.ComparePassword(input.Password, userData.Password) {
			return ctx.JSON(presenters.BadResponse(errors.New("wrong hutils")))
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["user_name"] = userData.Name
		claims["user_id"] = userData.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return ctx.JSON(presenters.SuccessRequest(map[string]string{"message": "Success login", "token": t}))
		//return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
	}

}
