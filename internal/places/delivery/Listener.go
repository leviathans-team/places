package delivery

import (
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal/places/delivery/hendlers"
)

func Hearing(app *fiber.App) {
	myGroup := app.Group("/place")
	myGroup.Get("", hendlers.GetItemById)
	myGroup.Get("/marketplace", hendlers.GetAll)
	myGroup.Get("/marketplace/type", hendlers.GetType)
	myGroup.Post("/marketplace/create", hendlers.PostNewItem)
	myGroup.Post("/marketplace/create_type", hendlers.PostNewType)
	myGroup.Put("/marketplace/update", hendlers.UpdateById)
	myGroup.Delete("marketplace/delete", hendlers.DeleteById)
}
