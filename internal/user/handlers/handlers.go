package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal"
	user "golang-pkg/internal/user/usecase"
	"log"
	"strconv"
	"time"
)

func SetupRoutesForAuth(app *fiber.App) {
	//commonUser := app.Group("")
	//
	//businessUser := app.Group("")

	admin := app.Group("/admin")
	admin.Put("/setAdmin", setAdmin)
	//admin.Put("/setAdmin", unSetAdmin)
	//admin.Put("/setAdmin", deleteProfile)
	//admin.Put("/setAdmin", deleteAdminProfile)

}

func setAdmin(ctx *fiber.Ctx) error {
	admLevel := ctx.Get("adminLevel", "")
	if admLevel != "3" {
		ctx.Status(401)
		return ctx.JSON(internal.HackError{
			Code:      401,
			Err:       errors.New("unauthorized admin"),
			Message:   "no rights",
			Timestamp: time.Now(),
		})
	}
	admLevelInt, err := strconv.ParseInt(admLevel, 10, 64)
	if err != nil {
		log.Print(err)
		ctx.Status(400)
		return ctx.JSON(internal.HackError{
			Code:      400,
			Err:       err,
			Message:   "incorrect value in header",
			Timestamp: time.Now(),
		})
	}
	hackErr := user.SetAdmin(admLevelInt)
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

//func unSetAdmin(ctx *fiber.Ctx) error {
//
//}
//
//func deleteProfile(ctx *fiber.Ctx) error {
//
//}
//
//func deleteAdminProfile(ctx *fiber.Ctx) error {
//
//}
