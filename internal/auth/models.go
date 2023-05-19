package auth

import "github.com/dgrijalva/jwt-go"

type UserForRegister struct {
	//Id         int    `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type BusinessUserForRegister struct { // Версия профиля для бизнеса
	//Id         int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserForLogin struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"userId"`
}
