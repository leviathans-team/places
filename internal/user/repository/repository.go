package userRepostiory

import (
	"errors"
	"github.com/lib/pq"
	"golang-pkg/internal"
	"log"
	"time"
)

func GetPlaceLandLord(userId int64) ([]int64, internal.HackError) {
	var slice pq.Int64Array
	err := internal.Tools.Connection.QueryRowx(`SELECT places FROM landlords WHERE user_id = $1`, userId).Scan(&slice)
	if err != nil {
		log.Print(err)
		return nil, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return slice, internal.HackError{}
}

func CreateNewPlace(userId, placeId int64) internal.HackError {
	var slice pq.Int64Array
	err := internal.Tools.Connection.QueryRowx(`SELECT places FROM landlords WHERE user_id = $1`, userId).Scan(&slice)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	slice = append(slice, placeId)
	_, err = internal.Tools.Connection.Exec(`UPDATE landlords SET places = $2 WHERE user_id=$1`, userId, placeId)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return internal.HackError{}
}

func SetAdmin(userId int64) internal.HackError {
	var isExist bool
	err := internal.Tools.Connection.QueryRowx(`SELECT EXISTS(select * from users_info where user_id=$1)`, userId).Scan(&isExist)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	if !isExist {
		log.Print(errors.New("invalid user id"))
		return internal.HackError{
			Code:      400,
			Err:       errors.New("invalid user id"),
			Timestamp: time.Now(),
		}
	}

	err = internal.Tools.Connection.QueryRowx(`select EXISTS(select from admins where user_id=$1)`, userId).Scan(&isExist)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	if isExist {
		log.Print(errors.New("user already admin"))
		return internal.HackError{
			Code:      400,
			Err:       errors.New("user already admin"),
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`INSERT INTO admins (user_id, admin_level) values ($1, 1)`, userId)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return internal.HackError{}
}

func GetAdminLevel(userId int64) (int64, internal.HackError) {
	var admlvl int64
	err := internal.Tools.Connection.QueryRowx(`select admin_level from admins where user_id=$1`, userId).Scan(&admlvl)
	if err != nil {
		log.Print(err)
		return 0, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return admlvl, internal.HackError{}
}

func IsLandlord(userId int64) (bool, internal.HackError) {
	var isExist bool
	err := internal.Tools.Connection.QueryRowx(`select exists(select * from landlords where user_id=$1)`, userId).Scan(&isExist)
	if err != nil {
		log.Print(err)
		return false, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return isExist, internal.HackError{}
}
