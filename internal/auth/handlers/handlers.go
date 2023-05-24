package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	"golang-pkg/internal/auth"
	"golang-pkg/internal/auth/usecase"
	"golang-pkg/middleware"
	"time"
)

func SetupRoutesForAuth(app *fiber.App) {
	api := app.Group("/user")
	api.Post("/login", login)
	api.Post("/register", register)
	api.Post("/businessRegister", landlordRegister)

	o2auth := api.Group("/o2auth")
	o2auth.Post("vk", loginWithVK)
	o2auth.Post("ok", loginWithOK)
	o2auth.Post("gos", loginWithGos)

	//test := app.Group("/test", middleware.UserIdentification)
	//test.Get("/123", testAuth)
	//test.Get("/124", testAuth2)

}

func login(ctx *fiber.Ctx) error {
	user := new(auth.UserForLogin)

	if err := ctx.BodyParser(user); err != nil {
		err := internal.HackError{
			Code:      400,
			Err:       err,
			Message:   "Can't parse json-body",
			Timestamp: time.Now(),
		}
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}

	errorsValidate := middleware.ValidateStruct(user)
	if errorsValidate != nil {
		ctx.Status(400)
		return ctx.JSON(errorsValidate)
	}

	token, err := usecase.SingIn(user)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(token)
}

func register(c *fiber.Ctx) error {
	user := new(auth.UserForRegister)
	if err := c.BodyParser(user); err != nil {
		err := internal.HackError{
			Code:      400,
			Err:       err,
			Message:   "Can't parse json-body",
			Timestamp: time.Now(),
		}
		c.Status(err.Code)
		return c.JSON(err)
	}

	errorsValidate := middleware.ValidateStruct(user)
	if errorsValidate != nil {
		c.Status(400)
		return c.JSON(errorsValidate)
	}
	err := usecase.SingUp(user)
	if err != nil {
		return c.JSON(err)
	}
	// ...

	return c.JSON(*user)
}

func landlordRegister(ctx *fiber.Ctx) error {
	user := new(auth.BusinessUserForRegister)
	if err := ctx.BodyParser(user); err != nil {
		err := internal.HackError{
			Code:      400,
			Err:       err,
			Message:   "Can't parse json-body",
			Timestamp: time.Now(),
		}
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}
	errorsValidate := middleware.ValidateStruct(user)
	if errorsValidate != nil {
		ctx.Status(400)
		return ctx.JSON(errorsValidate)
	}
	err := usecase.SingUpBusiness(user)
	if err != nil {
		return ctx.JSON(err)
	}
	// ...
	return ctx.JSON(*user)
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
