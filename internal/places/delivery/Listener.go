package delivery

import (
	"github.com/gofiber/fiber/v2"
	"golang-pkg/internal/places/delivery/hendlers"
	"golang-pkg/middleware"
)

func Hearing(app *fiber.App) {
	myGroup := app.Group("/place", middleware.UserIdentification)
	//myGroup := app.Group("/place")
	myGroup.Get("/chooseFilter", hendlers.GetAllFilters, middleware.UserIsExist)
	myGroup.Get("", hendlers.GetPlaces)
	myGroup.Get("/curent", hendlers.GetOnePlace, middleware.UserIsExist)
	myGroup.Post("/createFilter", hendlers.CreateFilter, middleware.UserIsExist)
	myGroup.Post("/createPlace", hendlers.CreatePlace, middleware.UserIsExist)
	myGroup.Delete("/delPlace", hendlers.DeletePlace, middleware.UserIsExist)
	myGroup.Delete("/delFilter", hendlers.DeleteFilter, middleware.UserIsExist)
	myGroup.Delete("/myOrders/cancelOrder", hendlers.CancelOrder, middleware.UserIsExist)
	app.Get("/myPlace", hendlers.GetMyPlaces, middleware.UserIdentification, middleware.UserIsExist)
	app.Get("/myOrders", hendlers.GetMyOrders, middleware.UserIdentification, middleware.UserIsExist)
	myGroup.Post("/createOrder", hendlers.CreateOrder, middleware.UserIsExist)
	myGroup.Put("/updatePlace", hendlers.UpdatePlace, middleware.UserIsExist)
	myGroup.Get("/searchPlace", hendlers.SearchPlace, middleware.UserIsExist)
	myGroup.Post("/curent/comments/createComment", hendlers.CreateComment, middleware.UserIsExist)
	myGroup.Get("/curent/comments", hendlers.GetComment, middleware.UserIsExist)

	myGroup.Post("/like", hendlers.CreateLike, middleware.UserIsExist)
	myGroup.Get("/getPlaceLikesCount", hendlers.GetPlaceLikeCount, middleware.UserIsExist)
	myGroup.Get("/isLiked", hendlers.IsLiked, middleware.UserIsExist)

	app.Get("/adminPlaces/placeForApproving", hendlers.GetNotApprovedPlace, middleware.UserIdentification, middleware.AdminIsExist)
	app.Get("/adminPlaces/placeForApproving/approve", hendlers.MakeApproved, middleware.UserIdentification, middleware.AdminIsExist)
}
