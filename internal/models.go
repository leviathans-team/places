package internal

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var Tools tools

type tools struct {
	Connection  *sqlx.DB
	Logger      *log.Logger
	AdminLogger *log.Logger
}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsAdmin    bool
	Password   string
}

type HackError struct {
	Code      int
	Err       error
	Message   string
	Timestamp time.Time
}

type UserHeaders struct {
	UserId     int64 `json:"userId"`
	IsLandLord bool  `json:"isLandLord"`
	AdminLevel int64 `json:"isAdmin"`
}
