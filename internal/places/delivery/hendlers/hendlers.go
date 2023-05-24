package hendlers

import (
	"github.com/gofiber/fiber/v2"
	models "golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"golang-pkg/internal/places/usecase"
	"log"
	"time"
)

// Возвращаю весь список фильров для выборки по ним
func GetAllFilters(ctx *fiber.Ctx) error {
	body, err := usecase.GetFilters()
	if err.Err != nil {
		log.Println(err)
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}
	return ctx.JSON(body)
}

// Создаю новый фильтр и возвращаю обновленный список фильров
func CreateFilter(ctx *fiber.Ctx) error {
	var body placeStruct.Filter
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]
	if err := ctx.BodyParser(&body); err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}

	result, err := usecase.CreateFilter(body, isAdmin)
	if err.Err != nil {
		log.Println(err)
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}
	return ctx.JSON(result)
}

// Создание нового места
func CreatePlace(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	isLandLord := headers["Islandlord"]
	userId := headers["userId"]
	var body placeStruct.Place
	err := ctx.BodyParser(&body)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}

	body, creationErr := usecase.CreatePlace(body, userId, isLandLord)
	if creationErr.Err != nil {
		log.Println(err)
		ctx.Status(creationErr.Code)
		return ctx.JSON(creationErr)
	}
	return ctx.JSON(body)

}

// Возвращает все места с учетом фильтра и с учетом даты
func GetPlaces(ctx *fiber.Ctx) error {
	var err error
	headers := ctx.GetReqHeaders()
	filter := headers["Filterid"]
	date := headers["Date"]
	page := headers["Page"]

	body, repError := usecase.GetPlaces(filter, date, page)

	if repError.Err != nil {
		log.Println(err)
		ctx.Status(repError.Code)
		return ctx.JSON(repError)
	}
	return ctx.Render("places", fiber.Map{
		"Places": body,
	})
}

// отдаю одно конкретное место по id
func GetOnePlace(ctx *fiber.Ctx) error {
	placeId := ctx.Query("placeId")
	body, repError := usecase.GetOnePlace(placeId)
	if repError.Err != nil {
		log.Println(repError.Err)
		ctx.Status(repError.Code)
		return ctx.JSON(repError)
	}
	return ctx.Render("place", fiber.Map{
		"Place": body,
	})
}

// удаление места. Доступно лендлордам и админам
func DeletePlace(ctx *fiber.Ctx) error {
	key := ctx.Query("placeId")
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]
	isLandLord := headers["Islandlord"]

	repErr := usecase.DeletePlace(key, isAdmin, isLandLord)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.SendStatus(200)
}

// удаление фильтра. Доступно только админам
func DeleteFilter(ctx *fiber.Ctx) error {
	filterid := ctx.Query("filterId")
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]

	repErr := usecase.DeleteFilter(filterid, isAdmin)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.SendStatus(200)
}

// отмена бронирования пользователем
func CancelOrder(ctx *fiber.Ctx) error {
	key := ctx.Query("orderId")
	headers := ctx.GetRespHeaders()
	userId := headers["Userid"]
	repErr := usecase.CancelOrder(key, userId)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.SendStatus(200)
}

func CreateOrder(ctx *fiber.Ctx) error {
	var body placeStruct.Calendar
	var result []placeStruct.Calendar
	err := ctx.BodyParser(&body)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	result, creationErr := usecase.CreateOrder(body)
	if creationErr.Err != nil {
		log.Println(creationErr.Err)
		ctx.Status(creationErr.Code)
		return ctx.JSON(creationErr)
	}
	return ctx.JSON(result)
}

func UpdatePlace(ctx *fiber.Ctx) error {
	var body placeStruct.Place
	err := ctx.BodyParser(&body)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	updateErr := usecase.UpdatePlace(body)
	if updateErr.Err != nil {
		log.Println(updateErr.Err)
		ctx.Status(updateErr.Code)
		return ctx.JSON(updateErr)
	}
	return ctx.SendStatus(200)
}

func SearchPlace(ctx *fiber.Ctx) error {
	key := ctx.Query("placeName")
	var result []placeStruct.Place

	result, searchErr := usecase.SearchPlace(key)
	if searchErr.Err != nil {
		log.Println(searchErr)
		ctx.Status(searchErr.Code)
		return ctx.JSON(searchErr)
	}
	return ctx.JSON(result)
}

// вывод собственных мест для лендлорда
func GetMyPlaces(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	userId := headers["Userid"]
	isLandLord := headers["Islandlord"]
	body, repErr := usecase.GetMyPlace(userId, isLandLord)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.JSON(body)
}

// возвращение всех бронирований пользователя
func GetMyOrders(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	userId := headers["Userid"]
	body, repErr := usecase.GetMyOrders(userId)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.JSON(body)
}

func CreateComment(ctx *fiber.Ctx) error {
	place := ctx.Query("placeId")
	headers := ctx.GetReqHeaders()
	userId := headers["Userid"]

	var body placeStruct.CommentMessage
	err := ctx.BodyParser(&body)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}

	result, repErr := usecase.CreateComment(userId, place, body)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.JSON(result)
}

func GetComment(ctx *fiber.Ctx) error {
	place := ctx.Query("placeId")
	result, repErr := usecase.GetComment(place)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.JSON(result)
}

func GetNotApprovedPlace(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]
	if isAdmin == "" {
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: errors.New("you can do this"), Timestamp: time.Now()})
	}
	body, creationErr := usecase.GetNotApprovedPlaces()
	if creationErr.Err != nil {
		ctx.Status(creationErr.Code)
		return ctx.JSON(creationErr)
	}
	return ctx.JSON(body)
}

func MakeApproved(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]
	if isAdmin == "" {
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: errors.New("you can do this"), Timestamp: time.Now()})
	}
	var body placeStruct.Approving
	err := ctx.BodyParser(&body)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	result, creationErr := usecase.MakeApproved(body.PlaceId)
	if creationErr.Err != nil {
		ctx.Status(creationErr.Code)
		return ctx.JSON(creationErr)
	}
	return ctx.JSON(result)
}
