package usecase

import (
	"errors"
	"golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"golang-pkg/internal/places/repository"
	user "golang-pkg/internal/user/usecase"
	"log"
	"strconv"
	"time"
)

func GetFilters() ([]placeStruct.Filter, *internal.HackError) {
	return repository.GetAllFilters()
}

func CreateFilter(body placeStruct.Filter, isAdmin string) ([]placeStruct.Filter, *internal.HackError) {
	if isAdmin == "" {
		return []placeStruct.Filter{}, &internal.HackError{Code: 400, Err: errors.New("you must be admin"), Timestamp: time.Now()}
	}
	return repository.CreateFilter(body)
}

func CreatePlace(body placeStruct.Place, user, isLandLord string) (placeStruct.Place, *internal.HackError) {
	if isLandLord == "false" {
		return placeStruct.Place{}, &internal.HackError{Code: 400, Err: errors.New("you must be landlord"), Timestamp: time.Now()}
	}

	userID, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		log.Println(err)
		return placeStruct.Place{}, &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}

	result, repErr := repository.CreatePlace(body)
	if repErr.Err != nil {
		return result, repErr
	}
	landUpdateError := user.CreateNewPlace(result.PlaceId, userID)
	if landUpdateError.Err != nil {
		return placeStruct.Place{}, landUpdateError
	}
	return result, repErr
}

func CreateOrder(body placeStruct.Calendar) ([]placeStruct.Calendar, *internal.HackError) {
	return repository.CreateOrder(body)
}

func UpdatePlace(body placeStruct.Place) *internal.HackError {
	return repository.UpdatePlace(body)
}

func SearchPlace(key string) ([]placeStruct.Place, *internal.HackError) {
	return repository.SearchPlace(key)
}

func GetPlaces(filter, dateH, page string) ([]placeStruct.Place, *internal.HackError) {
	date, err := time.Parse("2006-01-02 15:04:05", dateH)
	if err != nil {
		log.Println(err)
		return []placeStruct.Place{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}

	filterId := 0
	if filter != "" {
		filterId, err = strconv.Atoi(filter)
		if err != nil {
			log.Println(err)
			return []placeStruct.Place{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		}
	}
	pageNumber := 1
	if page != "" {
		pageNumber, err = strconv.Atoi(page)
		if err != nil {
			log.Println(err)
			return []placeStruct.Place{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
		}
	}
	return repository.GetPlaces(filterId, date, pageNumber)
}

func GetOnePlace(key string) (placeStruct.Place, *internal.HackError) {
	placeId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Println(err)
		return placeStruct.Place{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	return repository.GetOnePlace(placeId)
}
func DeletePlace(key, isAdmin, isLandLord string) *internal.HackError {

	placeId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Println(err)
		return &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	if isAdmin == "" || isLandLord == "false" {
		return &internal.HackError{Code: 400, Err: errors.New("you must be superuser"), Timestamp: time.Now()}
	}

	return repository.DeletePlace(placeId)
}

func DeleteFilter(key, isAdmin string) *internal.HackError {

	filterId, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Println(err)
		return &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	if isAdmin == "" {
		return &internal.HackError{Code: 400, Err: errors.New("you must be admin"), Timestamp: time.Now()}
	}

	return repository.DeleteFilter(filterId)
}

func CancelOrder(order string, user string) *internal.HackError {
	orderId, err := strconv.ParseInt(order, 10, 64)
	if err != nil {
		log.Println(err)
		return &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	userId, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		log.Println(err)
		return &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return repository.CancelOrder(orderId, userId)
}

func GetMyOrders(userId string) ([]placeStruct.Calendar, *internal.HackError) {
	user, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Calendar{}, &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return repository.GetMyOrders(user)
}

func GetMyPlace(userId, isLandLord string) ([]placeStruct.LandPlace, *internal.HackError) {

	landId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.LandPlace{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	if isLandLord == "false" {
		return []placeStruct.LandPlace{}, &internal.HackError{Code: 401, Err: errors.New("you must be landlord"), Timestamp: time.Now()}
	}

	placesId, err := user.GetPlacesLandlord(landId)
	if err != nil {
		log.Println(err)
		return []placeStruct.LandPlace{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	return repository.GetLandPlaces(placesId)
}

func CreateComment(user string, place string, body placeStruct.CommentMessage) ([]placeStruct.Comment, *internal.HackError) {
	placeId, err := strconv.ParseInt(place, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}

	userID, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	var repBody placeStruct.Comment
	repBody.Comment = body.Message
	repBody.UserId = userID
	repBody.PlaceId = placeId
	repBody.Mark = body.Mark
	return repository.CreateComment(repBody)

}

func GetComment(place string) ([]placeStruct.Comment, *internal.HackError) {
	placeId, err := strconv.ParseInt(place, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	return repository.GetComments(placeId)
}

func GetNotApprovedPlaces(isAdmin string) ([]placeStruct.Place, *internal.HackError) {
	if isAdmin == "" {
		return []placeStruct.Place{}, &internal.HackError{Code: 400, Err: errors.New("you must be admin"), Timestamp: time.Now()}
	}
	return repository.GetNotApprovedPlaces()
}

func MakeApproved(placeId int64, isAdmin string) (placeStruct.Place, *internal.HackError) {
	if isAdmin == "" {
		return placeStruct.Place{}, &internal.HackError{Code: 400, Err: errors.New("you must be admin"), Timestamp: time.Now()}
	}
	return repository.MakeApprove(placeId)
}

// TODO Добавить проверк уна существование пользователя
func CreateLike(place, user string) *internal.HackError {
	placeId, err := strconv.ParseInt(place, 10, 64)
	userId, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		log.Println(err)
		return &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	if _, placeExistErr := repository.GetOnePlace(placeId); placeExistErr.Err != nil {
		return &internal.HackError{Code: 404, Err: err, Timestamp: time.Now()}
	}

	return repository.CreateLike(placeId, userId)
}

func GetPlaceLikeCount(place string) (int64, *internal.HackError) {
	placeId, err := strconv.ParseInt(place, 10, 64)

	if err != nil {
		log.Println(err)
		return 0, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	if _, placeExistErr := repository.GetOnePlace(placeId); placeExistErr.Err != nil {
		return 0, &internal.HackError{Code: 404, Err: err, Timestamp: time.Now()}
	}

	return repository.GetPlaceLikeCount(placeId)
}

// TODO Добавить проверк уна существование пользователя
func IsLiked(place, user string) (bool, *internal.HackError) {
	userId, err := strconv.ParseInt(user, 10, 64)
	placeId, err := strconv.ParseInt(place, 10, 64)
	if err != nil {
		log.Println(err)
		return false, &internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}

	return repository.IsLiked(placeId, userId)
}
