package usecase

import (
	"golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"golang-pkg/internal/places/repository"
	"time"
)

func GetFilters() ([]placeStruct.Filter, internal.HackError) {
	return repository.GetAllFilters()
}

func CreateFilter(body placeStruct.Filter) ([]placeStruct.Filter, internal.HackError) {
	return repository.CreateFilter(body)
}

func CreatePlace(body placeStruct.Place) (placeStruct.Place, internal.HackError) {
	return repository.CreatePlace(body)
}

func GetPlaces(filterId int, date time.Time) ([]placeStruct.Place, internal.HackError) {
	return repository.GetPlaces(filterId, date, 1)
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
