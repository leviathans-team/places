package userRepostiory

import (
	"database/sql"
	"github.com/lib/pq"
	"golang-pkg/internal"
	"time"
)

func GetPlaceLandLord(userId int64) ([]int64, *internal.HackError) {
	var slice pq.Int64Array
	err := internal.Tools.Connection.QueryRowx(`SELECT places FROM landlords WHERE user_id = $1`, userId).Scan(&slice)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return nil, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return slice, nil
}

func CreateNewPlace(userId, placeId int64) *internal.HackError {
	var slice pq.Int64Array
	err := internal.Tools.Connection.QueryRowx(`SELECT places FROM landlords WHERE user_id = $1`, userId).Scan(&slice)
	if err != nil {
		if err == sql.ErrNoRows {
			slice = pq.Int64Array{}
		} else {
			internal.Tools.Logger.Print(err)
			return &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
	}

	slice = append(slice, placeId)
	_, err = internal.Tools.Connection.Exec(`UPDATE landlords SET places = $1 WHERE user_id = $2`, slice, userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return nil
}

func IsExistsOnAdminTable(userId int64) (bool, *internal.HackError) {
	var isExist bool
	err := internal.Tools.Connection.QueryRowx(`select EXISTS(select from admins where user_id=$1)`, userId).Scan(&isExist)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return false, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return isExist, nil
}

func IsExistsOnUsersTable(userId int64) (bool, *internal.HackError) {
	var isExist bool
	err := internal.Tools.Connection.QueryRowx(`SELECT EXISTS(select * from users_info where user_id=$1)`, userId).Scan(&isExist)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return false, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return isExist, nil
}

func IsExistsOnLandlordsTable(userId int64) (bool, *internal.HackError) {
	var isExist bool
	err := internal.Tools.Connection.QueryRowx(`SELECT EXISTS(select * from landlords where user_id=$1)`, userId).Scan(&isExist)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return false, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return isExist, nil
}

func SetAdmin(userId int64) *internal.HackError {
	_, err := internal.Tools.Connection.Exec(`INSERT INTO admins (user_id, admin_level) values ($1, 1)`, userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return nil
}

func UpAdminLevel(userId int64) *internal.HackError {
	_, err := internal.Tools.Connection.Exec(`UPDATE admins SET admin_level = admin_level+1 WHERE user_id=$1`, userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return nil
}

func GetAdminLevel(userId int64) (int64, *internal.HackError) {
	var admlvl int64
	err := internal.Tools.Connection.QueryRowx(`select admin_level from admins where user_id=$1`, userId).Scan(&admlvl)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return 0, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return admlvl, nil
}

func DeleteAdmin(userid int64) *internal.HackError {
	_, err := internal.Tools.Connection.Exec(`DELETE FROM admins WHERE user_id=$1`, userid)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return nil
}

func DeleteProfile(userId int64) *internal.HackError {
	_, err := internal.Tools.Connection.Exec(`DELETE from admins where user_id=$1`, userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`DELETE from landlords where user_id=$1`, userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`DELETE from users_login where login_id=$1`, userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`DELETE from users_info where user_id=$1`, userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return nil
}
