package usecase

import (
	"errors"
	"fmt"
	"github.com/mod/internal"
	"github.com/mod/internal/auth"
	"github.com/mod/internal/auth/repository"
	"log"
	"time"
)

func SingIn(user *auth.UserForLogin) internal.HackError {
	isExist, err := repository.ExistsUser(user.Login)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	if !isExist {
		return internal.HackError{
			Code:      404,
			Err:       errors.New("user not found"),
			Message:   "this email and phone is not found",
			Timestamp: time.Now(),
		}
	}

	if res, err := repository.TrySingIn(user.Login, user.Password); !res {
		if err.Err == nil {
			log.Print("failed login")
			err.Code = 400
			err.Err = errors.New("user not found")
			err.Message = "Check your login and password"
			err.Timestamp = time.Now()
		}
		return err
	} else {
		// return jwt
		fmt.Println("Success login!")
	}
	return internal.HackError{}
}

func SingUp(user *auth.UserForRegister) internal.HackError {
	existPhone, err := repository.ExistsUser(user.Phone)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	existEmail, err := repository.ExistsUser(user.Email)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	if existEmail || existPhone {
		log.Print("invaluable data")
		return internal.HackError{
			Code:      400,
			Err:       errors.New("invaluable data"),
			Message:   "the number or email is already taken",
			Timestamp: time.Now(),
		}
	}

	return repository.CreateUser(user)
}
