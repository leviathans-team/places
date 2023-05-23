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

func SetupRoutesForAuth(app *fiber.App) {
	//commonUser := app.Group("")
	//
	//businessUser := app.Group("")

	admin := app.Group("/admin", middleware.UserIdentification)
	admin.Put("/setAdmin/:userId", setAdmin)
	admin.Put("/promotionAdmin", promotionAdmin)

	admin.Put("/setAdmin", unSetAdmin)
	admin.Put("/setAdmin", deleteProfile)
	admin.Put("/setAdmin", deleteAdminProfile)

}

func setAdmin(ctx *fiber.Ctx) error {

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

	//admLevelInt, err := strconv.ParseInt(admLevel, 10, 64)
	//if err != nil {
	//	log.Print(err)
	//	ctx.Status(400)
	//	return ctx.JSON(internal.HackError{
	//		Code:      400,
	//		Err:       err,
	//		Message:   "incorrect value in header",
	//		Timestamp: time.Now(),
	//	})
	//}
	hackErr := user.SetAdmin(userIdInt)
	if hackErr.Err != nil {
		ctx.Status(hackErr.Code)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       err,
			Message:   "incorrect value in header",
			Timestamp: time.Now(),
		})
	}
	return nil
}

func unSetAdmin(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func deleteProfile(ctx *fiber.Ctx) error {
	return ctx.Next()

}

func deleteAdminProfile(ctx *fiber.Ctx) error {
	return ctx.Next()

}

func promotionAdmin(ctx *fiber.Ctx) error {
	//err := user.PromotionAdmin(1)
	return ctx.Next()
	//return ctx.Next()
}
