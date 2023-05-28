package userHandlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	_ "golang-pkg/internal/user"
	user "golang-pkg/internal/user/usecase"
	"golang-pkg/middleware"
	"strconv"
	"time"
)

func UserPanel(app *fiber.App) {
	user := app.Group("/user", middleware.UserIdentification, middleware.UserIsExist)
	user.Get("/info", getUser)

	landlord := app.Group("landlord", middleware.UserIdentification, middleware.UserIsExist)
	landlord.Get("/info", getLandlord)

	admin := app.Group("/admin", middleware.UserIdentification, middleware.AdminIsExist)
	admin.Put("/setAdmin/id/:userId", setAdmin)
	admin.Put("/promotionAdmin/id/:userId", promotionAdmin)
	admin.Put("/unSetAdmin/id/:userId", unSetAdmin)
	admin.Put("/deleteProfile/id/:userId", deleteProfile)
	admin.Put("/deleteAdminProfile/id/:userId", deleteAdminProfile)
}

// @Summary Получение информации о пользователе
// @Tags User
// @Security ApiKeyAuth
// @Description Получение информации о пользователе (ФИО, номер, почта)
// @ID getUser
// @Param userId header string true "ИД пользователя"
// @Produce  json
// @Success 200 {object} userModels.User
// @Failure 400 {object} internal.HackError
// @Failure 401 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /user/info [get]
// Возвращаю userModels.User если  все успешно
func getUser(ctx *fiber.Ctx) error {

	userId := ctx.GetRespHeader("userId", "")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		hackErr := internal.HackError{
			Code:      400,
			Err:       err,
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	if body, hackErr := user.GetUserData(userIdInt); hackErr != nil {
		ctx.Status(hackErr.Code)
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	} else {
		return ctx.JSON(body)
	}
}

// @Summary Получение информации о пользователе
// @Tags 	Landlord
// @Security ApiKeyAuth
// @Description Получение информации о пользователе (ФИО, номер, почта, должность, его места, ЮР лицо, ИНН)
// @ID getLandlord
// @Param userId header string true "ИД пользователя"
// @Produce  json
// @Success 200 {object} userModels.Landlord
// @Failure 400 {object} internal.HackError
// @Failure 401 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /landlord/info [get]
// Возвращаю userModels.Landlord если  все успешно
func getLandlord(ctx *fiber.Ctx) error {
	userId := ctx.GetRespHeader("userId", "")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		hackErr := internal.HackError{
			Code:      400,
			Err:       err,
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	flag, hackErr := user.IsLandlord(userIdInt)
	if err != nil {
		ctx.Status(hackErr.Code)
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	if !flag {
		ctx.Status(401)
		hackErr = &internal.HackError{
			Code:      400,
			Err:       errors.New("user is not landlord"),
			Message:   err.Error(),
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	body, hackErr := user.GetLandlordData(userIdInt)
	if hackErr != nil {
		ctx.Status(hackErr.Code)
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	return ctx.JSON(body)
}

// @Summary Назначение администратора
// @Tags Admin
// @Security ApiKeyAuth
// @Description Авторизировать пользователя
// @ID setAdmin
// @Param userId header string true "ИД администратора"
// @Param adminLevel header int true "Уровень администратора, где 0-не админ, 1 - админ младщего звена, 3 - старший админ"
// @Param userid path int true "Ид пользователя, которого назначают администратором"
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
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		hackErr := internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	hackErr := user.SetAdmin(adminIdInt, userIdInt)
	if hackErr != nil {
		ctx.Status(hackErr.Code)
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return nil
}

// @Summary Снятие с поста администратора
// @Tags Admin
// @Security ApiKeyAuth
// @Description Снятие с поста администратора
// @ID unSetAdmin
// @Param userId header string true "ИД администратора"
// @Param adminLevel header int true "Уровень администратора, где 0-не админ, 1 - админ младщего звена, 3 - старший админ"
// @Param userid path int true "Ид пользователя, которого назначают администратором"
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
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		hackErr := internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	hackErr := user.UnSetAdmin(adminIdInt, userIdInt)
	if hackErr != nil {
		ctx.Status(hackErr.Code)
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return nil
}

// @Summary Удаление аккаунта
// @Tags Admin
// @Security ApiKeyAuth
// @Description Удаление аккаунта не админиистратора
// @ID deleteProfile
// @Param userId header string true "ИД администратора"
// @Param adminLevel header int true "Уровень администратора, где 0-не админ, 1 - админ младщего звена, 3 - старший админ"
// @Param userid path int true "Ид пользователя, которого назначают администратором"
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
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	admLevel := ctx.GetRespHeader("Adminlevel", "")
	if !(admLevel == "2" || admLevel == "3") {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		hackErr := internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "uncorrected params",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	hackErr := user.DeleteProfile(adminIdInt, userIdInt)
	if hackErr != nil {
		ctx.Status(hackErr.Code)
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return nil
}

// @Summary Удаление аккаунта администратора
// @Tags Admin
// @Security ApiKeyAuth
// @Description Удаление аккаунта не админиистратора
// @ID deleteAdminProfile
// @Param userId header string true "ИД администратора"
// @Param adminLevel header int true "Уровень администратора, где 0-не админ, 1 - админ младщего звена, 3 - старший админ"
// @Param userid path int true "Ид пользователя, которого назначают администратором"
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
// @Param userId header string true "ИД администратора"
// @Param adminLevel header int true "Уровень администратора, где 0-не админ, 1 - админ младщего звена, 3 - старший админ"
// @Param userid path int true "Ид пользователя, которого назначают администратором"
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
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		internal.Tools.Logger.Print(errors.New("unauthorized admin"))
		ctx.Status(401)
		hackErr := internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	userId := ctx.Params("userid", "-1")
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		ctx.Status(400)
		hackErr := internal.HackError{
			Code:      400,
			Err:       errors.New("uncorrected params"),
			Message:   "uncorrected params",
			Timestamp: time.Now(),
		}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	hackErr := user.PromotionAdmin(adminIdInt, userIdInt)
	if hackErr != nil {
		ctx.Status(hackErr.Code)
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return nil
	//return ctx.Next()
}
