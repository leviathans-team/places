package hendlers

import (
	"github.com/gofiber/fiber/v2"
	models "golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"golang-pkg/internal/places/usecase"
	"log"
	"strconv"
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
// Нужна валидация на ADmin
func CreateFilter(ctx *fiber.Ctx) error {
	var body placeStruct.Filter
	if err := ctx.BodyParser(body); err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	result, err := usecase.CreateFilter(body)
	if err.Err != nil {
		log.Println(err)
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}
	return ctx.JSON(result)
}

// Создание нового места
// Надо докрутить валидацию на landlord и передачу Алмазу
func CreatePlace(ctx *fiber.Ctx) error {
	var body placeStruct.Place
	err := ctx.BodyParser(&body)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	body, creationErr := usecase.CreatePlace(body)
	if creationErr.Err != nil {
		log.Println(err)
		ctx.Status(creationErr.Code)
		return ctx.JSON(creationErr)
	}
	return ctx.JSON(body)

}

// Возвращает все места с учетом фильтра и с учетом даты
func GetPlaces(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	filter := headers["filterId"]
	date := headers["date"]
	page := headers["page"]

	body, repError := usecase.GetPlaces(filter, date, page)
	if repError.Err != nil {
		log.Println(repError.Err)
		ctx.Status(repError.Code)
		return ctx.JSON(repError)
	}
	return ctx.JSON(body)
}

// отдаю одно конкретное место по id
func GetOnePlace(ctx *fiber.Ctx) error {
	key := ctx.Query("placeId")
	placeId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	body, repError := usecase.GetOnePlace(placeId)
	if repError.Err != nil {
		log.Println(err)
		ctx.Status(repError.Code)
		return ctx.JSON(repError)
	}
	return ctx.JSON(body)
}

func DeletePlace(ctx *fiber.Ctx) error {
	key := ctx.Query("placeId")
	placeId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	repErr := usecase.DeletePlace(placeId)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.SendStatus(200)
}

func DeleteFilter(ctx *fiber.Ctx) error {
	key := ctx.Query("filterId")
	filterId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Println(err)
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	repErr := usecase.DeleteFilter(filterId)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.SendStatus(200)
}

func CancelOrder(ctx *fiber.Ctx) error {
	key := ctx.Query("orderId")
	repErr := usecase.CancelOrder(key)
	if repErr.Err != nil {
		log.Println(repErr)
		ctx.Status(repErr.Code)
		return ctx.JSON(repErr)
	}
	return ctx.SendStatus(200)
}
