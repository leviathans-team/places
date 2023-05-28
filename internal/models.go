package internal

import (
	"encoding/json"
	"fmt"
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

func (e *HackError) Error() string {
	return fmt.Sprintf("Status: %d\n"+
		"Error: %s\n"+
		"Message: %s\n"+
		"Timestamp: %s\n",
		e.Code, e.Err.Error(), e.Message, e.Timestamp.String())
}

func (reqErr *HackError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code      int    `json:"code"`
		Err       string `json:"err"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
	}{
		Code:      reqErr.Code,
		Err:       reqErr.Err.Error(),
		Message:   reqErr.Message,
		Timestamp: reqErr.Timestamp.String(),
	})
}
