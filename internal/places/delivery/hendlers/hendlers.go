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
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}
	return ctx.JSON(body)
}

// Создаю новый фильтр и возвращаю обновленный список фильров
func CreateFilter(ctx *fiber.Ctx) error {
	var body placeStruct.Filter
	if err := ctx.BodyParser(body); err != nil {
		ctx.Status(400)
		return ctx.JSON(models.HackError{Code: 400, Err: err, Timestamp: time.Now()})
	}
	result, err := usecase.CreateFilter(body)
	if err.Err != nil {
		ctx.Status(err.Code)
		return ctx.JSON(err)
	}
	return ctx.JSON(result)
}

func CreatePlace(ctx *fiber.Ctx) error {

}

func GetPlaces(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	key := headers["filterId"]
	filterId, err := strconv.Atoi(key)
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(usecase.GetPlaces(filterId))
}

func GetOnePlace(ctx *fiber.Ctx) error {
	key := ctx.Query("placeId")
	placeId, err := strconv.Atoi(key)
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(usecase.GetOnePlace(placeId))
}
