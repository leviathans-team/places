package repository

import (
	"crypto/md5"
	"errors"
	"fmt"
	"golang-pkg/internal"
	"golang-pkg/internal/auth"
	"time"
)

var salt = "eJNJhefjkJKfkjherf"

func ExistsUser(login string) (bool, error) {
	var isExists bool = false
	err := internal.Tools.Connection.QueryRowx(`select exists(select * from users_info where email=$1 or phone=$1)`, login).Scan(&isExists)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return false, err
	}
	return isExists, nil
}

func CreateUser(user *auth.UserForRegister) *internal.HackError {
	var userId int64
	passwordHash := md5.New()
	passwordHash.Write([]byte(user.Password))
	passwordHash.Sum([]byte(salt))
	stringPasswordHash := fmt.Sprintf("%x", passwordHash)

	err := internal.Tools.Connection.QueryRowx(`INSERT into users_info (name, surname, patronymic, email, phone) 
values ($1, $2, $3, $4, $5) returning user_id`, user.Name, user.Surname, user.Patronymic, user.Email, user.Phone).Scan(&userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`insert into users_login values ($1, $2)`, userId, stringPasswordHash)
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

func CreateBusinessUser(user *auth.BusinessUserForRegister) *internal.HackError {
	var userId int64
	passwordHash := md5.New()
	passwordHash.Write([]byte(user.Password))
	passwordHash.Sum([]byte(salt))
	stringPasswordHash := fmt.Sprintf("%x", passwordHash)

	err := internal.Tools.Connection.QueryRowx(`INSERT into users_info (name, surname, patronymic, email, phone) 
values ($1, $2, $3, $4, $5) returning user_id`, user.Name, user.Surname, user.Patronymic, user.Email, user.Phone).Scan(&userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`insert into users_login values ($1, $2)`, userId, stringPasswordHash)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`insert into landlords (user_id, post, legal_entity, inn)
values ($1,$2,$3,$4)`, userId, user.Post, user.LegalEntity, user.INN)
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

func TrySingIn(login, password string) (int64, *internal.HackError) {
	result := false
	var userId int64
	passwordHash := md5.New()
	passwordHash.Write([]byte(password))
	passwordHash.Sum([]byte(salt))
	stringPasswordHash := fmt.Sprintf("%x", passwordHash)

	err := internal.Tools.Connection.QueryRowx(`select user_id from users_info where (email=$1 or phone=$1)`, login).Scan(&userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return -1, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	err = internal.Tools.Connection.QueryRowx(`select exists(select * from users_login where login_id=$1 and password_hash=$2)`, userId, stringPasswordHash).Scan(&result)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return -1, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	if !result {
		internal.Tools.Logger.Print(errors.New("failed login"))
		return -1, &internal.HackError{
			Code:      401,
			Err:       errors.New("failed login"),
			Message:   "login or password incorrect",
			Timestamp: time.Now(),
		}
	}
	return userId, nil
}
