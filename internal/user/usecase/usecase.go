package user

import (
	"golang-pkg/internal"
	userRepostiory "golang-pkg/internal/user/repository"
)

func SetAdmin(userId int64) internal.HackError {
	err := userRepostiory.SetAdmin(userId)
	return err
}

func CreateNewPlace(placeId, userId int64) internal.HackError {
	err := userRepostiory.CreateNewPlace(placeId, userId)
	return err
}

func GetPlacesLandlord(userId int64) ([]int64, internal.HackError) {
	places, err := userRepostiory.GetPlaceLandLord(userId)
	return places, err
}

func IsAdmin(userId int64) (int64, internal.HackError) {
	return userRepostiory.GetAdminLevel(userId)
}

func IsLandlord(userId int64) (bool, internal.HackError) {
	return userRepostiory.IsLandlord(userId)
}

func PromotionAdmin(userId int64) internal.HackError {
	return internal.HackError{}
}
