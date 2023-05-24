package delivery

import (
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal/places/delivery/hendlers"
)

func Hearing(app *fiber.App) {
	//myGroup := app.Group("/place", middleware.UserIdentification)
	myGroup := app.Group("/place")
	myGroup.Get("/chooseFilter", hendlers.GetAllFilters)
	myGroup.Get("", hendlers.GetPlaces)
	myGroup.Get("/curent", hendlers.GetOnePlace)
	myGroup.Post("/createFilter", hendlers.CreateFilter)
	myGroup.Post("/createPlace", hendlers.CreatePlace)
	myGroup.Delete("/delPlace", hendlers.DeletePlace)
	myGroup.Delete("/delFilter", hendlers.DeleteFilter)
	myGroup.Delete("/myOrders/cancelOrder", hendlers.CancelOrder)
	app.Get("/myPlace", hendlers.GetMyPlaces)
	app.Get("/myOrders", hendlers.GetMyOrders)
	myGroup.Delete("/myOrders/cancelOrder", hendlers.CancelOrder)
	myGroup.Post("/createOrder", hendlers.CreateOrder)
	myGroup.Put("/updatePlace", hendlers.UpdatePlace)
	myGroup.Get("/searchPlace", hendlers.SearchPlace)
	myGroup.Post("/curent/comments/createComment", hendlers.CreateComment)
	myGroup.Get("/curent/comments", hendlers.GetComment)

	app.Get("/adminPlaces/placeForApproving", hendlers.GetNotApprovedPlace)
	app.Get("/adminPlaces/placeForApproving/approve", hendlers.MakeApproved)
}
