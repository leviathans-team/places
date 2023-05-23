package user

import (
	"errors"
	"golang-pkg/internal"
	userRepostiory "golang-pkg/internal/user/repository"
	"log"
	"time"
)

func CreateNewPlace(placeId, userId int64) internal.HackError {
	err := userRepostiory.CreateNewPlace(placeId, userId)
	return err
}

func GetPlacesLandlord(userId int64) ([]int64, internal.HackError) {
	places, err := userRepostiory.GetPlaceLandLord(userId)
	return places, err
}

func IsAdmin(userId int64) (int64, internal.HackError) {
	isExist, hackErr := userRepostiory.IsExistsOnAdminTable(userId)
	if hackErr.Err != nil {
		return 0, hackErr
	}
	if isExist {
		return userRepostiory.GetAdminLevel(userId)
	}
	return 0, internal.HackError{}
}

func IsLandlord(userId int64) (bool, internal.HackError) {
	return userRepostiory.IsExistsOnLandlordsTable(userId)
}

func SetAdmin(adminId, userId int64) internal.HackError {
	var isExist bool
	isExist, hackErr := userRepostiory.IsExistsOnUsersTable(userId)
	if hackErr.Err != nil {
		return hackErr
	}

	if !isExist {
		log.Print(errors.New("invalid user id"))
		return internal.HackError{
			Code:      400,
			Err:       errors.New("invalid user id"),
			Timestamp: time.Now(),
		}
	}

	isExist, hackErr = userRepostiory.IsExistsOnAdminTable(userId)
	if hackErr.Err != nil {
		return hackErr
	}

	if isExist {
		log.Print(errors.New("user already admin"))
		return internal.HackError{
			Code:      400,
			Err:       errors.New("user already admin"),
			Timestamp: time.Now(),
		}
	}
	err := userRepostiory.SetAdmin(userId)
	if err.Err != nil {
		return err
	}

	internal.Tools.AdminLogger.Printf("Admin (id: %ld) granted administrator rights to the user (id: %ld)", adminId, userId)
	return internal.HackError{}
}

func PromotionAdmin(adminId, userId int64) internal.HackError {
	isExist, err := userRepostiory.IsExistsOnAdminTable(userId)
	if err.Err != nil {
		return err
	}
	if !isExist {
		log.Print(errors.New("invalid userId"))
		return internal.HackError{
			Code:      400,
			Err:       errors.New("invalid userId"),
			Message:   "This Id does not belong to the administrator",
			Timestamp: time.Now(),
		}
	}

	adminLevel, err := userRepostiory.GetAdminLevel(userId)
	if err.Err != nil {
		return err
	}
	if adminLevel == 3 {
		log.Print(errors.New("admin have max level"))
		return internal.HackError{
			Code:      400,
			Err:       errors.New("admin have max level"),
			Message:   "admin have max level",
			Timestamp: time.Now(),
		}
	}

	err = userRepostiory.UpAdminLevel(userId)
	if err.Err != nil {
		return err
	}

	internal.Tools.AdminLogger.Printf("Admin (id: %d) has increased the access level for the user (id: %d) to level %d",
		adminId, userId, adminLevel+1)
	return internal.HackError{}
}

//func UnSetAdmin(adminId, userId int64) internal.HackError {
//
//}
