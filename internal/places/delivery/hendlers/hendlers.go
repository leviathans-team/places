package hendlers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	models "golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"golang-pkg/internal/places/usecase"
	"time"
)

// @Summary Получить полных список фильтров
// @Security ApiKeyAuth
// @Tags places
// @Description Получить полных список фильтров
// @ID GetAllFilters
// @Produce  json
// @Success 200 {array} []placeStruct.Filter{}
// @Failure 400 {object} internal.HackError
// @Failure 404 {object} internal.HackError
// @Failure 500 {object} internal.HackError
// @Router /place/chooseFilter [get]
// Возвращаю весь список фильров для выборки по ним
func GetAllFilters(ctx *fiber.Ctx) error {
	body, err := usecase.GetFilters()
	if err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(err.Code)
		jErr, _ := err.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(body)
}

// @Summary Создание нового фильтра
// @Security ApiKeyAuth
// @Tags places
// @Description Создание нового фильтра и возврат обновленного списка фильтров
// @ID CreateFilter
// @Accept  json
// @Produce  json
// @Param        Isadmin   header      string  false  "Является админом(true/false)"
// @Param        Filter   body    placeStruct.Filter{}    true  "Json для создания фильтра"
// @Success 200 {array} placeStruct.Filter{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/createFilter [post]
func CreateFilter(ctx *fiber.Ctx) error {
	var body placeStruct.Filter
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]
	if err := ctx.BodyParser(&body); err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(400)
		hackErr := models.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	result, err := usecase.CreateFilter(body, isAdmin)
	if err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(err.Code)
		jErr, _ := err.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}

// @Summary Создание нового места
//
// @Security ApiKeyAuth
// @Tags places
// @Description Создание нового места и возврат этого места
// @ID CreatePlace
// @Accept  json
// @Produce  json
// @Param        Islandlord   header      string  false  "Является лэндлордом(true/false)"
// @Param        Userid   header      string  false  "ID пользователя"
// @Param        Place   body    placeStruct.Place{}    true  "Json для создания места"
// @Success 200 {object} placeStruct.Place{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/createPlace [post]
func CreatePlace(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	isLandLord := headers["Islandlord"]
	userId := headers["Userid"]
	var body placeStruct.Place
	err := ctx.BodyParser(&body)
	if err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(400)
		hackErr := models.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}

	body, creationErr := usecase.CreatePlace(body, userId, isLandLord)
	if creationErr != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(creationErr.Code)
		jErr, _ := creationErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(body)

}

// @Summary Получение всех мест
// @Tags places
// @Description Получение всех мест и отображение их на странице
// @ID GetPlaces
// @Accept  json
// @Produce  html
// @Param        Filterid   header      string  false  "ID фильтра"
// @Param        Date   header      string  false  "Дата бронирования"
// @Param        Page   header      string  false  "Страница для пагинации"
// @Success 200 {object} []placeStruct.Place{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router / [get]
func GetPlaces(ctx *fiber.Ctx) error {
	var err error
	headers := ctx.GetReqHeaders()
	filter := headers["Filterid"]
	date := headers["Date"]
	page := headers["Page"]

	body, repError := usecase.GetPlaces(filter, date, page)

	if repError != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(repError.Code)
		jErr, _ := repError.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.Render("places", fiber.Map{
		"Places": body,
	})
}

// @Summary Получение конкретного места
// @Security ApiKeyAuth
// @Tags places
// @Description получение конкретного места по id
// @ID GetOnePlace
// @Produce  html
// @Param        placeId   query      string  false  "ID места"
// @Success 200 {object} placeStruct.Place{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/curent [get]
func GetOnePlace(ctx *fiber.Ctx) error {
	placeId := ctx.Query("placeId")
	body, repError := usecase.GetOnePlace(placeId)
	if repError != nil {
		models.Tools.Logger.Println(repError)
		ctx.Status(repError.Code)
		jErr, _ := repError.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.Render("place", fiber.Map{
		"Place": body,
	})
}

// @Summary Удаление конкретного места
// @Security ApiKeyAuth
// @Tags places
// @Description удаление места. Доступно лендлордам и админам
// @ID DeletePlace
// @Produce  json
// @Param        placeId   query      string  false  "ID места"
// @Param        Islandlord   header      string  false  "Является лэндлордом(true/false)"
// @Param        Isadmin   header      string  false  "Является админом(true/false)"
// @Success 200 {string} OK!
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/delPlace [delete]
func DeletePlace(ctx *fiber.Ctx) error {
	key := ctx.Query("placeId")
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]
	isLandLord := headers["Islandlord"]

	repErr := usecase.DeletePlace(key, isAdmin, isLandLord)
	if repErr != nil {
		models.Tools.Logger.Println(repErr)
		ctx.Status(repErr.Code)
		jErr, _ := repErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.SendStatus(200)
}

// @Summary Удаление фильтра
// @Security ApiKeyAuth
// @Tags places
// @Description удаление фильтра по ID. Доступно только админам
// @ID DeleteFilter
// @Produce  json
// @Param        filterId   query      string  false  "ID фильтра"
// @Param        Isadmin   header      string  false  "Является админом(true/false)"
// @Success 200 {string} OK!
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/delFilter [delete]
func DeleteFilter(ctx *fiber.Ctx) error {
	filterid := ctx.Query("filterId")
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]

	repErr := usecase.DeleteFilter(filterid, isAdmin)
	if repErr != nil {
		models.Tools.Logger.Println(repErr)
		ctx.Status(repErr.Code)
		jErr, _ := repErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.SendStatus(200)
}

// @Summary Отмена бронирования пользователем
// @Security ApiKeyAuth
// @Tags places
// @Description отмена бронирования пользователем по orderId и Userid
// @ID CancelOrder
// @Produce  json
// @Param        orderId   query      string  false  "ID фильтра"
// @Param        Userid   header      string  false  "ID пользователя"
// @Success 200 {string} OK!
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/myOrders/cancelOrder [delete]
func CancelOrder(ctx *fiber.Ctx) error {
	key := ctx.Query("orderId")
	headers := ctx.GetRespHeaders()
	userId := headers["Userid"]
	repErr := usecase.CancelOrder(key, userId)
	if repErr != nil {
		models.Tools.Logger.Println(repErr)
		ctx.Status(repErr.Code)
		jErr, _ := repErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.SendStatus(200)
}

// @Summary Создание бронирования
// @Security ApiKeyAuth
// @Tags places
// @Description Создание бронирования и получение списка заказов у конкретного места
// @ID CreateOrder
// @Produce  json
// @Param        Order   body    placeStruct.Filter{}    true  "Json для создания брони"
// @Success 200 {array} []placeStruct.Calendar{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/createOrder [post]
func CreateOrder(ctx *fiber.Ctx) error {
	var body placeStruct.Calendar
	var result []placeStruct.Calendar
	err := ctx.BodyParser(&body)
	if err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(400)
		hackErr := models.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	result, creationErr := usecase.CreateOrder(body)
	if creationErr != nil {
		models.Tools.Logger.Println(creationErr)
		ctx.Status(creationErr.Code)
		jErr, _ := creationErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}

// @Summary Обновление параметров места
// @Security ApiKeyAuth
// @Tags places
// @Description Обновление параметров места
// @ID UpdatePlace
// @Produce  json
// @Param        Order   body    placeStruct.Place{}    true  "Json для обновления места"
// @Success 200 {string} OK!
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/updatePlace [put]
func UpdatePlace(ctx *fiber.Ctx) error {
	var body placeStruct.Place
	err := ctx.BodyParser(&body)
	if err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(400)
		hackErr := models.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	updateErr := usecase.UpdatePlace(body)
	if updateErr != nil {
		models.Tools.Logger.Println(updateErr)
		ctx.Status(updateErr.Code)
		jErr, _ := updateErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.SendStatus(200)
}

// @Summary Поиск мест
// @Security ApiKeyAuth
// @Tags places
// @Description Поиск мест по его названию
// @ID SearchPlace
// @Produce  json
// @Param      placeName  query    string    true  "Название места"
// @Success 200 {array} []placeStruct.Place{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/searchPlace [get]
func SearchPlace(ctx *fiber.Ctx) error {
	key := ctx.Query("placeName")
	var result []placeStruct.Place

	result, searchErr := usecase.SearchPlace(key)
	if searchErr != nil {
		models.Tools.Logger.Println(searchErr)
		ctx.Status(searchErr.Code)
		jErr, _ := searchErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}

// @Summary Вывод свох мест
// @Security ApiKeyAuth
// @Tags Ownner
// @Description вывод собственных мест для лендлорда
// @ID GetMyPlaces
// @Produce  json
// @Param      Userid  header    string    true  "ID пользователя"
// @Param      Islandlord   header      string  false  "Является лэндлордом(true/false)"
// @Success 200 {array} []placeStruct.LandPlace
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/myPlace [get]
func GetMyPlaces(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	userId := headers["Userid"]
	isLandLord := headers["Islandlord"]
	body, repErr := usecase.GetMyPlace(userId, isLandLord)
	if repErr != nil {
		models.Tools.Logger.Println(repErr)
		ctx.Status(repErr.Code)
		jErr, _ := repErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(body)
}

// @Summary Вывод свох бронирований
// @Security ApiKeyAuth
// @Tags Ownner
// @Description возвращение всех бронирований пользователя
// @ID GetMyOrders
// @Produce  json
// @Param      Userid  header    string    true  "ID пользователя"
// @Success 200 {array} []placeStruct.Calendar{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /myOrders [get]
func GetMyOrders(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	userId := headers["Userid"]
	body, repErr := usecase.GetMyOrders(userId)
	if repErr != nil {
		models.Tools.Logger.Println(repErr)
		ctx.Status(repErr.Code)
		jErr, _ := repErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(body)
}

// @Summary Оставить отзыв
// @Security ApiKeyAuth
// @Tags places
// @Description Оставить комментарий и поставить оценку
// @ID CreateComment
// @Produce  json
// @Param      placeId  query    string    true  "ID места"
// @Param      Userid  header    string    true  "ID пользователя"
// @Param      Comment  body    placeStruct.CommentMessage{}    true  "ID пользователя"
// @Success 200 {array} []placeStruct.Comment{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/curent/comments/createComment [post]
func CreateComment(ctx *fiber.Ctx) error {
	place := ctx.Query("placeId")
	headers := ctx.GetReqHeaders()
	userId := headers["Userid"]

	var body placeStruct.CommentMessage
	err := ctx.BodyParser(&body)
	if err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(400)
		hackErr := models.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.JSON(jErr)
	}

	result, repErr := usecase.CreateComment(userId, place, body)
	if repErr != nil {
		models.Tools.Logger.Println(repErr)
		ctx.Status(repErr.Code)
		jErr, _ := repErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}

// @Summary Посмотреть оценки
// @Security ApiKeyAuth
// @Tags places
// @Description Получить оценку
// @ID GetComment
// @Produce  json
// @Param      placeId  query    string    true  "ID места"
// @Success 200 {array} []placeStruct.Comment{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /place/curent/comments [get]
func GetComment(ctx *fiber.Ctx) error {
	place := ctx.Query("placeId")
	result, repErr := usecase.GetComment(place)
	if repErr != nil {
		models.Tools.Logger.Println(repErr)
		ctx.Status(repErr.Code)
		jErr, _ := repErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}

// @Summary Получить не подтверждённые места
// @Security ApiKeyAuth
// @Tags Admin
// @Description Получить не подтверждённые места
// @ID GetNotApprovedPlace
// @Produce  json
// @Param   Isadmin   header      string  false  "Является админом(true/false)"
// @Success 200 {array} []placeStruct.Place{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /adminPlaces/placeForApproving [get]
func GetNotApprovedPlace(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]

	body, creationErr := usecase.GetNotApprovedPlaces(isAdmin)
	if creationErr != nil {
		ctx.Status(creationErr.Code)
		jErr, _ := creationErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(body)
}

// @Summary Подтвердить место
// @Security ApiKeyAuth
// @Tags Admin
// @Description Подтвердить место
// @ID MakeApproved
// @Produce  json
// @Param   Isadmin   header      string  false  "Является админом(true/false)"
// @Param   Place   body      placeStruct.Approving  false  "Подтверждение места"
// @Success 200 {array} placeStruct.Place{}
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /adminPlaces/placeForApproving [put]
func MakeApproved(ctx *fiber.Ctx) error {
	headers := ctx.GetRespHeaders()
	isAdmin := headers["Isadmin"]
	var body placeStruct.Approving
	err := ctx.BodyParser(&body)
	placeId := body.PlaceId
	if err != nil {
		models.Tools.Logger.Println(err)
		ctx.Status(400)
		hackErr := models.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		jErr, _ := hackErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	result, creationErr := usecase.MakeApproved(placeId, isAdmin)
	if creationErr != nil {
		ctx.Status(creationErr.Code)
		jErr, _ := creationErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}

// @Summary Добавить место в избранное
// @Security ApiKeyAuth
// @Tags places
// @Description Добавить место в избранное
// @ID CreateLike
// @Produce  json
// @Param   placeId   query      string  true  "ID места"
// @Param   userId   query      string  true  "ID пользователя"
// @Success 200 {string} OK!
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /places/like [post]
func CreateLike(ctx *fiber.Ctx) error {
	placeId := ctx.Query("placeId")
	userId := ctx.Query("userId")
	createErr := usecase.CreateLike(placeId, userId)
	if createErr != nil {
		models.Tools.Logger.Println(createErr)
		ctx.Status(createErr.Code)
		jErr, _ := createErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.SendStatus(200)
}

// @Summary Получить количество лайков у места
// @Security ApiKeyAuth
// @Tags places
// @Description Добавить место в избранное
// @ID GetPlaceLikeCount
// @Produce  json
// @Param   placeId   query   string  true  "ID места"
// @Success 200 {int} 123
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /places/getPlaceLikesCount [get]
func GetPlaceLikeCount(ctx *fiber.Ctx) error {
	placeId := ctx.Query("placeId")
	result, getErr := usecase.GetPlaceLikeCount(placeId)
	if getErr != nil {
		models.Tools.Logger.Println(getErr)
		ctx.Status(getErr.Code)
		jErr, _ := getErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}

// @Summary Проверить лайкнуто или нет
// @Security ApiKeyAuth
// @Tags places
// @Description Проверить лайкнуто или нете
// @ID IsLiked
// @Produce  json
// @Param   placeId   query      string  true  "ID места"
// @Param   userId   query      string  true  "ID пользоватя"
// @Success 200 {string} OK
// @Failure 400 {string} internal.HackError.Err
// @Failure 404 {string} internal.HackError.Err
// @Failure 500 {string} internal.HackError.Err
// @Router /places/isLiked [get]
func IsLiked(ctx *fiber.Ctx) error {
	placeId := ctx.Query("placeId")
	userId := ctx.Query("userId")
	result, getErr := usecase.IsLiked(placeId, userId)
	if getErr != nil {
		models.Tools.Logger.Println(getErr)
		ctx.Status(getErr.Code)
		jErr, _ := getErr.MarshalJSON()
		return ctx.Send(jErr)
	}
	return ctx.JSON(result)
}
