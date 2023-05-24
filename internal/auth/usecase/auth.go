package usecase

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang-pkg/internal"
	"golang-pkg/internal/auth"
	"golang-pkg/internal/auth/repository"
	user "golang-pkg/internal/user/usecase"
	"time"
)

var (
	tokenTTL  = 10 * time.Hour
	singInKey = "rhnJHfjhrgjke8Nihe843Hgherekrr3e4fgrf"
)

func SingUp(user *auth.UserForRegister) *internal.HackError {
	existPhone, err := repository.ExistsUser(user.Phone)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	existEmail, err := repository.ExistsUser(user.Email)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	if existEmail || existPhone {
		internal.Tools.Logger.Print("invaluable data")
		return &internal.HackError{
			Code:      400,
			Err:       errors.New("invaluable data"),
			Message:   "the number or email is already taken",
			Timestamp: time.Now(),
		}
	}

	return repository.CreateUser(user)
}

func SingIn(user *auth.UserForLogin) (string, *internal.HackError) {
	isExist, err := repository.ExistsUser(user.Login)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return "", &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	if !isExist {
		return "", &internal.HackError{
			Code:      404,
			Err:       errors.New("user not found"),
			Message:   "this email and phone is not found",
			Timestamp: time.Now(),
		}
	}
	userId, hackErr := repository.TrySingIn(user.Login, user.Password)
	if hackErr != nil {
		return "", hackErr
	}

	token, hackErr := GenerateToken(userId)

	if hackErr != nil {
		return "", hackErr
	}

	// return jwt
	fmt.Println("Success login!")

	return token, nil
}

func SingUpBusiness(user *auth.BusinessUserForRegister) *internal.HackError {
	isExist, err := repository.ExistsUser(user.Phone)
	if err != nil {
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	if isExist {
		return &internal.HackError{
			Code:      400,
			Err:       errors.New("invaluable data"),
			Message:   "the number or email is already taken",
			Timestamp: time.Now(),
		}
	}

	isExist, err = repository.ExistsUser(user.Email)
	if err != nil {
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	if isExist {
		return &internal.HackError{
			Code:      400,
			Err:       errors.New("invaluable data"),
			Message:   "the number or email is already taken",
			Timestamp: time.Now(),
		}
	}

	hackErr := repository.CreateBusinessUser(user)
	if hackErr != nil {
		return hackErr
	}
	return nil
}

func GenerateToken(userId int64) (string, *internal.HackError) {
	isLadLord, err := user.IsLandlord(userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return "", err
	}
	lvlAdmin, err := user.IsAdmin(userId)
	if err != nil {
		internal.Tools.Logger.Print(err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), //Токен будет жить 10 часов
			IssuedAt:  time.Now().Unix(),
		},
		UserId:     userId,
		IsLandLord: isLadLord,
		AdminLevel: lvlAdmin,
	})

	tokenString, e := token.SignedString([]byte(singInKey))
	if e != nil {
		internal.Tools.Logger.Fatal("FIX 110 line in repository")
	}

	return tokenString, nil
}

func ParseToken(accessToken string) (*internal.UserHeaders, *internal.HackError) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(singInKey), nil
	})
	if err != nil {
		internal.Tools.Logger.Print("invalid signing method")
		return nil, &internal.HackError{
			Code:      400,
			Err:       errors.New("invalid signing method"),
			Timestamp: time.Now(),
		}
	}
	claims, ok := token.Claims.(*auth.TokenClaims)
	if !ok {
		return nil, &internal.HackError{
			Code:      400,
			Err:       errors.New("token claims are not a type"),
			Timestamp: time.Now(),
		}
	}

	return &internal.UserHeaders{
		UserId:     claims.UserId,
		IsLandLord: claims.IsLandLord,
		AdminLevel: claims.AdminLevel,
	}, nil
}
