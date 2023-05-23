package usecase

import (
	"golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"golang-pkg/internal/places/repository"
	"log"
	"strconv"
	"time"
)

func GetFilters() ([]placeStruct.Filter, internal.HackError) {
	return repository.GetAllFilters()
}

func CreateFilter(body placeStruct.Filter) ([]placeStruct.Filter, internal.HackError) {
	return repository.CreateFilter(body)
}

func CreatePlace(body placeStruct.Place, user string) (placeStruct.Place, internal.HackError) {
	userID, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		log.Println(err)
		return placeStruct.Place{}, internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
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

func CreateOrder(body placeStruct.Calendar) ([]placeStruct.Calendar, internal.HackError) {
	return repository.CreateOrder(body)
}

func UpdatePlace(body placeStruct.Place) internal.HackError {
	return repository.UpdatePlace(body)
}

func SearchPlace(key string) ([]placeStruct.Place, internal.HackError) {
	return repository.SearchPlace(key)
}

func GetPlaces(filterId int, date time.Time, page int) ([]placeStruct.Place, internal.HackError) {
	return repository.GetPlaces(filterId, date, page)
}

func GetOnePlace(placeId int64) (placeStruct.Place, internal.HackError) {
	return repository.GetOnePlace(placeId)
}

func DeletePlace(placeId int64) internal.HackError {
	return repository.DeletePlace(placeId)
}

func DeleteFilter(filterId int64) internal.HackError {
	return repository.DeleteFilter(filterId)
}

func CancelOrder(order string, user string) internal.HackError {
	orderId, err := strconv.ParseInt(order, 10, 64)
	if err != nil {
		log.Println(err)
		return internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	userId, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		log.Println(err)
		return internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return repository.CancelOrder(orderId, userId)
}

func GetMyOrders(userId string) ([]placeStruct.Calendar, internal.HackError) {
	user, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Calendar{}, internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return repository.GetMyOrders(user)
}

func GetMyPlace(userId int64) ([]placeStruct.LandPlace, internal.HackError) {
	placesId, err := user.GetPlacesLandlord(userId)
	if err.Err != nil {
		return []placeStruct.LandPlace{}, err
	}
	return repository.GetLandPlaces(placesId)
}

func CreateComment(user string, place string, body placeStruct.CommentMessage) ([]placeStruct.Comment, internal.HackError) {
	placeId, err := strconv.ParseInt(place, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}

	userID, err := strconv.ParseInt(user, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	var repBody placeStruct.Comment
	repBody.Comment = body.Message
	repBody.UserId = userID
	repBody.PlaceId = placeId
	repBody.Mark = body.Mark
	return repository.CreateComment(repBody)

}

func GetComment(place string) ([]placeStruct.Comment, internal.HackError) {
	placeId, err := strconv.ParseInt(place, 10, 64)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, internal.HackError{Code: 400, Err: err, Timestamp: time.Now()}
	}
	return repository.GetComments(placeId)
}
