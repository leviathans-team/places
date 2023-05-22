package delivery

import (
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal/places/delivery/hendlers"
	"golang-pkg/middleware"
)

func Hearing(app *fiber.App) {
	myGroup := app.Group("/place", middleware.UserIdentification)
	myGroup.Get("/chooseFilter", hendlers.GetAllFilters)
	myGroup.Get("", hendlers.GetPlaces)
	myGroup.Get("/curent", hendlers.GetOnePlace)
	myGroup.Post("/createFilter", hendlers.CreateFilter)
	myGroup.Post("/createPlace", hendlers.CreatePlace)
	myGroup.Delete("/delPlace", hendlers.DeletePlace)
	myGroup.Delete("/delFilter", hendlers.DeleteFilter)
	myGroup.Delete("/cancelOrder")
	myGroup.Post("/createOrder")
	myGroup.Put("/updatePlace")
	myGroup.Get("/searchPlace")
	myGroup.Get("/myPlace")

}
