package usecase

import (
	"golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"golang-pkg/internal/places/repository"
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
