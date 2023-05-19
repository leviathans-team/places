package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	"golang-pkg/internal/auth"
	"golang-pkg/internal/auth/usecase"

	"time"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/user")
	api.Post("/login", login)
	api.Post("/register", register)

	o2auth := api.Group("/o2auth")
	o2auth.Post("vk", loginWithVK)
	o2auth.Post("ok", loginWithOK)
	o2auth.Post("gos", loginWithGos)

}

func login(ctx *fiber.Ctx) error {
	user := new(auth.UserForLogin)

	if err := ctx.BodyParser(user); err != nil {
		err := internal.HackError{
			Code:      400,
			Err:       err,
			Message:   "",
			Timestamp: time.Now(),
		}
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}

	errors := auth.ValidateStruct(user)
	if errors != nil {
		ctx.Status(400)
		return ctx.JSON(errors)
	}

	err := usecase.SingIn(user)
	if err.Err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(*user)
}

func register(c *fiber.Ctx) error {
	user := new(auth.UserForRegister)
	if err := c.BodyParser(user); err != nil {
		err := internal.HackError{
			Code:      400,
			Err:       err,
			Message:   "",
			Timestamp: time.Now(),
		}
		c.Status(err.Code)
		return c.JSON(err)
	}

	errors := auth.ValidateStruct(user)
	if errors != nil {
		c.Status(400)
		return c.JSON(errors)
	}
	err := usecase.SingUp(user)
	if err.Err != nil {
		return c.JSON(err)
	}
	// ...

	return c.JSON(*user)
}

func loginWithGos(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return errors.New("in progress")
}

func loginWithOK(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return errors.New("in progress")
}

func loginWithVK(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return errors.New("in progress")
}
