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
	app.Get("getUserInfo", middleware.UserIdentification)
	api := app.Group("/auth")
	api.Post("/login", login)
	api.Post("/register", register)
	api.Post("/businessRegister", landlordRegister)

	o2auth := api.Group("/o2auth")
	o2auth.Post("vk", loginWithVK)
	o2auth.Post("Tinkoff", loginWithTinkoff)
	o2auth.Post("Sber", loginWithSber)
	o2auth.Post("gos", loginWithGos)

	//test := app.Group("/test", middleware.UserIdentification)
	//test.Get("/123", testAuth)
	//test.Get("/124", testAuth2)

}

// @Summary Авторизация
// @Tags auth
// @Description Авторизировать пользователя
// @ID login
// @Param        loginDate   body    auth.UserForLogin{}    true  "Данные для входа"
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /auth/login [post]
// Возвращаю jwt-token, который храним в себе данные об авторизации
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
		jErr, _ := err.MarshalJSON()
		return ctx.Send(jErr)
		//return ctx.JSON(err)
	}

	errorsValidate := middleware.ValidateStruct(user)
	if errorsValidate != nil {
		ctx.Status(400)
		return ctx.JSON(errorsValidate)
	}

	token, err := usecase.SingIn(user)
	if err != nil {
		jErr, _ := err.MarshalJSON()
		return ctx.Send(jErr)
		//return ctx.JSON(err)
	}
	return ctx.JSON(token)
}

// @Summary Регистрация
// @Tags auth
// @Description Регистрация пользователя
// @ID register
// @Param        regDate   body    auth.UserForRegister{}    true  "Данные для регистрации"
// @Produce  json
// @Success 200 {object} auth.UserForRegister
// @Failure 400 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /auth/register [post]
// Возвращаю в случае успеха данные, введеные при регистрации
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
		jErr, _ := err.MarshalJSON()
		return c.Send(jErr)
	}

	errorsValidate := middleware.ValidateStruct(user)
	if errorsValidate != nil {
		c.Status(400)
		return c.JSON(errorsValidate)
	}
	err := usecase.SingUp(user)
	if err != nil {
		jErr, _ := err.MarshalJSON()
		return c.Send(jErr)
	}
	// ...

	return c.JSON(*user)
}

// @Summary Регистрация арендодателя
// @Tags auth
// @Description Регистрация пользователя
// @ID landlordRegister
// @Param        regDate   body    auth.BusinessUserForRegister{}    true  "Данные для регистрации"
// @Produce  json
// @Success 200 {object} auth.BusinessUserForRegister
// @Failure 400 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /auth/businessRegister [post]
// Возвращаю в случае успеха данные, введеные при регистрации
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
		jErr, _ := err.MarshalJSON()
		return ctx.Send(jErr)
	}
	errorsValidate := middleware.ValidateStruct(user)
	if errorsValidate != nil {
		ctx.Status(400)
		return ctx.JSON(errorsValidate)
	}
	err := usecase.SingUpBusiness(user)
	if err != nil {
		jErr, _ := err.MarshalJSON()
		return ctx.Send(jErr)
	}
	// ...
	return ctx.JSON(*user)
}

// @Summary Вход через госуслуги
// @Tags auth
// @Description Регистрация пользователя -> Заглушка
// @ID loginWithGos
// @Failure 404 {object} error
// @Router /auth/o2auth/gos [post]
// Заглушка
func loginWithGos(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return errors.New("in progress")
}

// @Summary Вход через Sber
// @Tags auth
// @Description Регистрация пользователя -> Заглушка
// @ID loginWithSber
// @Failure 404 {object} error
// @Router /auth/o2auth/Sber [post]
// Заглушка
func loginWithSber(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return errors.New("in progress")
}

// @Summary Вход через VK
// @Tags auth
// @Description Регистрация пользователя -> Заглушка
// @ID loginWithVK
// @Failure 404 {object} error
// @Router /auth/o2auth/svk [post]
// Заглушка
func loginWithVK(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return errors.New("in progress")
}

// @Summary Вход через Tinkoff
// @Tags auth
// @Description Регистрация пользователя -> Заглушка
// @ID loginWithTinkoff
// @Failure 404 {object} error
// @Router /auth/o2auth/Tinkoff [post]
// Заглушка
func loginWithTinkoff(ctx *fiber.Ctx) error {
	ctx.Status(404)
	return errors.New("in progress")
}
