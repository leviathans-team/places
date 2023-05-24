package userHandlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	user "golang-pkg/internal/user/usecase"
	"golang-pkg/middleware"
	"log"
	"strconv"
	"time"
)

func UserPanel(app *fiber.App) {
	//commonUser := app.Group("")
	//
	//businessUser := app.Group("")
	//app.Post("test/isAdmin", test)

	admin := app.Group("/admin", middleware.UserIdentification, middleware.AdminIsExist)
	admin.Put("/setAdmin/id/:userId", setAdmin)
	admin.Put("/promotionAdmin/id/:userId", promotionAdmin)

	admin.Put("/unSetAdmin/id/:userId", unSetAdmin)
	admin.Put("/deleteProfile/id/:userId", deleteProfile)
	admin.Put("/deleteAdminProfile/id/:userId", deleteAdminProfile)

}

//func test(ctx *fiber.Ctx) error {
//	level, err := user.IsAdmin(1)
//	fmt.Println(level)
//	fmt.Println(err)
//	return ctx.Next()
//}

func setAdmin(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		log.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		log.Print(errors.New("unauthorized admin"))
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
			Code:      401,
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

func unSetAdmin(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		log.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		log.Print(errors.New("unauthorized admin"))
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
			Code:      401,
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

func deleteProfile(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		log.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("Adminlevel", "")
	if !(admLevel == "2" || admLevel == "3") {
		log.Print(errors.New("unauthorized admin"))
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
			Code:      401,
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

func promotionAdmin(ctx *fiber.Ctx) error {
	adminId := ctx.GetRespHeader("userId", "")
	adminIdInt, err := strconv.ParseInt(adminId, 10, 64)
	if err != nil {
		log.Print(errors.New("invalid header userId"))
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       errors.New("invalid header userId"),
			Timestamp: time.Now(),
		})
	}
	admLevel := ctx.GetRespHeader("adminLevel", "")
	if admLevel != "3" {
		log.Print(errors.New("unauthorized admin"))
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
			Code:      401,
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
