package hendlers

//
//import (
//	"github.com/gofiber/fiber/v2"
//	models "golang-pkg/internal"
//	"golang-pkg/internal/usecase"
//	"log"
//	"strconv"
//)
//
//func GetAllFilters(ctx *fiber.Ctx) error {
//	return ctx.JSON(usecase.GetFilters())
//}
//
//func CreateFilter(ctx *fiber.Ctx) error {
//	body := new(models.Filter)
//	if err := ctx.BodyParser(body); err != nil {
//		return err
//	}
//
//}
//
//func GetPlaces(ctx *fiber.Ctx) error {
//	headers := ctx.GetReqHeaders()
//	key := headers["filterId"]
//	filterId, err := strconv.Atoi(key)
//	if err != nil {
//		log.Println(err)
//	}
//	return ctx.JSON(usecase.GetPlaces(filterId))
//}
//
//func GetOnePlace(ctx *fiber.Ctx) error {
//	key := ctx.Query("placeId")
//	placeId, err := strconv.Atoi(key)
//	if err != nil {
//		log.Println(err)
//	}
//	return ctx.JSON(usecase.GetOnePlace(placeId))
//}
//
////
////func GetItemById(ctx *fiber.Ctx) error {
////	key := ctx.Query("id")
////	id, err := strconv.Atoi(key)
////	if err != nil {
////		log.Print(err)
////	}
////	var body models.Filter
////	body = usecase.GetItem(id)
////	return ctx.JSON(body)
////}
////
////func GetType(ctx *fiber.Ctx) error {
////	types := usecase.GetTypes()
////	return ctx.JSON(types)
////}
////
////func GetAll(ctx *fiber.Ctx) error {
////	headers := ctx.GetReqHeaders()
////	startId := headers["From_id"]
////	itemType := headers["Item_type"]
////	if startId == "" {
////		if itemType == "" {
////			return ctx.JSON(usecase.GetItems(0, 0))
////		} else {
////			types, err := strconv.Atoi(itemType)
////			if err != nil {
////				log.Print(err)
////			}
////			return ctx.JSON(usecase.GetItems(0, types))
////
////		}
////	} else {
////		intId, err := strconv.Atoi(startId)
////		if err != nil {
////			log.Print(err)
////		}
////		if itemType == "" {
////			return ctx.JSON(usecase.GetItems(intId, 0))
////		} else {
////			types, err := strconv.Atoi(itemType)
////			if err != nil {
////				log.Print(err)
////			}
////			return ctx.JSON(usecase.GetItems(intId, types))
////
////		}
////	}
////	return ctx.SendStatus(400)
////
////}
////
////func PostNewItem(ctx *fiber.Ctx) error {
////	body := new(marketplace.Items)
////	body.Default()
////	if err := ctx.BodyParser(body); err != nil {
////		return err
////	}
////	usecase.CreateNewItem(body)
////	return nil
////}
////
////func UpdateById(ctx *fiber.Ctx) error {
////	body := new(marketplace.Items)
////	if err := ctx.BodyParser(body); err != nil {
////		return err
////	}
////
////	return ctx.JSON(usecase.UpdateItem(*body))
////}
////
////func DeleteById(ctx *fiber.Ctx) error {
////	headers := ctx.GetReqHeaders()
////	id := headers["Id"]
////	intId, err := strconv.Atoi(id)
////	if err != nil {
////		log.Print(err)
////	}
////	usecase.DeleteItem(intId)
////	return nil
////}
////
////func PostNewType(ctx *fiber.Ctx) error {
////	headers := ctx.GetReqHeaders()
////	newtype := headers["Newtype"]
////	if newtype == "" {
////		return nil
////	}
////	usecase.CreateType(newtype)
////	return nil
////}
