package auth

import "github.com/dgrijalva/jwt-go"

type UserForRegister struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type BusinessUserForRegister struct { // Версия профиля для бизнеса
	Name        string `json:"name" validate:"required"`
	Surname     string `json:"surname" validate:"required"`
	Patronymic  string `json:"patronymic" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Post        string `json:"post" validate:"required"`
	LegalEntity string `json:"legalEntity" validate:"required"`
	INN         string `json:"inn" validate:"required"`
}

type UserForLogin struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId     int64 `json:"userId"`
	IsLandLord bool  `json:"isLandLord"`
	IsAdmin    bool  `json:"isAdmin"`
}
