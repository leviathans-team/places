package usecase

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang-pkg/internal"
	"golang-pkg/internal/auth"
	"golang-pkg/internal/auth/repository"

	"log"
	"time"
)

var (
	tokenTTL  = 10 * time.Hour
	singInKey = "rhnJHfjhrgjke8Nihe843Hgherekrr3e4fgrf"
)

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

func SingIn(user *auth.UserForLogin) (string, internal.HackError) {
	isExist, err := repository.ExistsUser(user.Login)
	if err != nil {
		log.Print(err)
		return "", internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	if !isExist {
		return "", internal.HackError{
			Code:      404,
			Err:       errors.New("user not found"),
			Message:   "this email and phone is not found",
			Timestamp: time.Now(),
		}
	}
	userId, res := repository.TrySingIn(user.Login, user.Password)
	if res.Err != nil {
		return "", res
	}

	token, hackErr := GenerateToken(userId)

	if hackErr.Err != nil {
		return "", hackErr
	}

	// return jwt
	fmt.Println("Success login!")

	return token, internal.HackError{}
}

func SingUpBusiness(user *auth.BusinessUserForRegister) internal.HackError {
	isExist, err := repository.ExistsUser(user.Phone)
	if err != nil {
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	if isExist {
		return internal.HackError{
			Code:      400,
			Err:       errors.New("invaluable data"),
			Message:   "the number or email is already taken",
			Timestamp: time.Now(),
		}
	}

	isExist, err = repository.ExistsUser(user.Email)
	if err != nil {
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	if isExist {
		return internal.HackError{
			Code:      400,
			Err:       errors.New("invaluable data"),
			Message:   "the number or email is already taken",
			Timestamp: time.Now(),
		}
	}

	hackErr := repository.CreateBusinessUser(user)
	if hackErr.Err != nil {
		return hackErr
	}
	return internal.HackError{}
}

func GenerateToken(userId int64) (string, internal.HackError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), //Токен будет жить 10 часов
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})

	tokenString, e := token.SignedString([]byte(singInKey))
	if e != nil {
		log.Fatal("FIX 110 line in repository")
	}

	return tokenString, internal.HackError{}
}

func ParseToken(accessToken string) (*internal.UserHeaders, internal.HackError) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(singInKey), nil
	})
	if err != nil {
		log.Print("invalid signing method")
		return nil, internal.HackError{
			Code:      400,
			Err:       errors.New("invalid signing method"),
			Timestamp: time.Now(),
		}
	}
	claims, ok := token.Claims.(*auth.TokenClaims)
	if !ok {
		return nil, internal.HackError{
			Code:      400,
			Err:       errors.New("token claims are not a type"),
			Timestamp: time.Now(),
		}
	}

	return &internal.UserHeaders{
		UserId:     claims.UserId,
		IsLandLord: claims.IsLandLord,
		AdminLevel: claims.AdminLevel,
	}, internal.HackError{}
}
