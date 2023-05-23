package models

import (
	"github.com/lib/pq"
	"time"
)

type Filter struct {
	FilterId   int64  `json:"filterId"`
	FilterName string `json:"filterName"`
}

type Place struct {
	PlaceId         int64          `json:"placeId"`
	PlaceName       string         `json:"placeName"`
	FilterId        int64          `json:"filterId"`
	PlaceAddress    string         `json:"placeAddress"`
	WorkingTime     string         `json:"workingTime"`
	TelephoneNumber string         `json:"telephoneNumber"`
	Email           string         `json:"email"`
	Site            string         `json:"site"`
	PlaceServices   string         `json:"placeServices"`
	TotalSquare     float64        `json:"totalSquare"`
	WorkingSquare   float64        `json:"workingSquare"`
	CommonObjects   string         `json:"commonObjects"`
	Equipment       string         `json:"equipment"`
	RentersCount    int64          `json:"rentersCount"`
	Meta            pq.StringArray `json:"meta"`
}

type Calendar struct {
	BookId   int64     `json:"bookId"`
	PlaceId  int64     `json:"placeId"`
	TimeFrom time.Time `json:"timeFrom"`
	TimeTo   time.Time `json:"timeTo"`
	UserId   int64     `json:"userId"`
}

type LandPlace struct {
	Place
	Calendar
}
