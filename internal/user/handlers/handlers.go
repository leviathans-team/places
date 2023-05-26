package userHandlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	user "golang-pkg/internal/user/usecase"
	"golang-pkg/middleware"
	"strconv"
	"time"
)

func UserPanel(app *fiber.App) {
	admin := app.Group("/admin", middleware.UserIdentification, middleware.AdminIsExist)
	admin.Put("/setAdmin/id/:userId", setAdmin)
	admin.Put("/promotionAdmin/id/:userId", promotionAdmin)

	admin.Put("/unSetAdmin/id/:userId", unSetAdmin)
	admin.Put("/deleteProfile/id/:userId", deleteProfile)
	admin.Put("/deleteAdminProfile/id/:userId", deleteAdminProfile)
}

// @Summary Назначение администратора
// @Tags Admin
// @Security ApiKeyAuth
// @Description Авторизировать пользователя
// @ID setAdmin
// @Params input header userId true "ИД администратора"
// @Params input header adminLevel true "Уровень администратора"
// @Params input path userid true "Ид пользователя, которого назначают администратором"
// @Produce  json
// @Success 200 {object} nil
// @Failure 400 {object} internal.HackError
// @Failure 401 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /admin/setAdmin/id/:userId [put]
// Возвращаю nil если  все успешно
func setAdmin(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		internal.Tools.Logger.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		})
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "",
			Timestamp: time.Now(),
		})
	}
	hackErr := user.SetAdmin(adminIdInt, userIdInt)
	if hackErr != nil {
		ctx.Status(hackErr.Code)
		return ctx.JSON(hackErr)
	}
	return nil
}

// @Summary Снятие с поста администратора
// @Tags Admin
// @Security ApiKeyAuth
// @Description Снятие с поста администратора
// @ID unSetAdmin
// @Params input header userId true "ИД администратора"
// @Params input header adminLevel true "Уровень администратора"
// @Params input path userid true "Ид пользователя, которого снимают с поста администратора"
// @Produce  json
// @Success 200 {object} nil
// @Failure 400 {object} internal.HackError
// @Failure 401 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /admin/usSetAdmin/id/:userId [put]
// Возвращаю nil если  все успешно
func unSetAdmin(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		internal.Tools.Logger.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		})
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "",
			Timestamp: time.Now(),
		})
	}

	hackErr := user.UnSetAdmin(adminIdInt, userIdInt)
	if hackErr != nil {
		ctx.Status(hackErr.Code)
		return ctx.JSON(hackErr)
	}
	return nil
}

// @Summary Удаление аккаунта
// @Tags Admin
// @Security ApiKeyAuth
// @Description Удаление аккаунта не админиистратора
// @ID deleteProfile
// @Params input header userId true "ИД администратора"
// @Params input header adminLevel true "Уровень администратора"
// @Params input path userid true "Ид пользователя которому удаляют аккаунт"
// @Produce  json
// @Success 200 {object} nil
// @Failure 400 {object} internal.HackError
// @Failure 401 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /admin/deleteProfile/id/:userId [put]
// Возвращаю nil если  все успешно
func deleteProfile(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		internal.Tools.Logger.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("Adminlevel", "")
	if !(admLevel == "2" || admLevel == "3") {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		})
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "uncorrected params",
			Timestamp: time.Now(),
		})
	}

	hackErr := user.DeleteProfile(adminIdInt, userIdInt)
	if hackErr != nil {
		return ctx.JSON(hackErr)
	}
	return nil
}

// @Summary Удаление аккаунта администратора
// @Tags Admin
// @Security ApiKeyAuth
// @Description Удаление аккаунта не админиистратора
// @ID deleteAdminProfile
// @Params input header userId true "ИД администратора"
// @Params input header adminLevel true "Уровень администратора"
// @Params input path userid true "Ид пользователя которому удаляют аккаунт"
// @Produce  json
// @Success 200 {object} nil
// @Failure 400 {object} internal.HackError
// @Failure 401 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /admin/deleteAdminProfile/id/:userId [put]
// Возвращаю nil если  все успешно
func deleteAdminProfile(ctx *fiber.Ctx) error {
	err := unSetAdmin(ctx)
	if err != nil {
		return ctx.JSON(err)
	}
	err = deleteProfile(ctx)
	if err != nil {
		return ctx.JSON(err)
	}
	return nil
}

// @Summary Повышение уровня администратора
// @Tags Admin
// @Security ApiKeyAuth
// @Description Повышение уровня администратора
// @ID promotionAdmin
// @Params input header userId true "ИД администратора"
// @Params input header adminLevel true "Уровень администратора"
// @Params input path userid true "Ид пользователя, которому повысят уровень администратора"
// @Produce  json
// @Success 200 {object} nil
// @Failure 400 {object} internal.HackError
// @Failure 401 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /admin/promotionAdmin/id/:userId [put]
// Возвращаю nil если  все успешно
func promotionAdmin(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		internal.Tools.Logger.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		})
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "uncorrected params",
			Timestamp: time.Now(),
		})
	}

	hackErr := user.PromotionAdmin(adminIdInt, userIdInt)
	if hackErr != nil {
		return ctx.JSON(hackErr)
	}
	return nil
	//return ctx.Next()
}
