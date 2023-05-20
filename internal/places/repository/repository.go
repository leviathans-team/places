package repository

import (
	"golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"log"
	"time"
)

func CreateFilter(body placeStruct.Filter) ([]placeStruct.Filter, internal.HackError) {
	var filterId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO filters
	(filterName)
	VALUES ($1)`, body.FilterName).Scan(&filterId)
	if err != nil {
		return []placeStruct.Filter{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return GetAllFilters()
}

func GetAllFilters() ([]placeStruct.Filter, internal.HackError) {
	var result []placeStruct.Filter
	err := internal.Tools.Connection.Get(&result, `SELECT * FROM filters`)
	if err != nil {
		return []placeStruct.Filter{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return result, internal.HackError{}
}

func CreatePlace(body placeStruct.Place) (placeStruct.Place, internal.HackError) {
	var placeId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO places
    (placeName, filterId, placeAddress, workingTime, telephoneNumber, email, site, placeServices, totalSquare, workingSquare,
     commonObjects, equipment, rentersCount, meta) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning placeId`, body.PlaceName,
		body.FilterId, body.PlaceAddress, body.WorkingTime, body.TelephoneNumber, body.Email, body.Site, body.PlaceServices,
		body.TotalSquare, body.WorkingSquare, body.CommonObjects, body.Equipment, body.RentersCount, body.Meta).Scan(&placeId)
	if err != nil {
		log.Println(err)
		return placeStruct.Place{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return GetOnePlace(placeId)
}

func GetOnePlace(placeId int64) (placeStruct.Place, internal.HackError) {
	var body placeStruct.Place
	err := internal.Tools.Connection.Get(&body, `SELECT * FROM places WHERE placeId = $1`, placeId)
	if err != nil {
		return placeStruct.Place{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return body, internal.HackError{}
}

func CreateCalendarNote(body placeStruct.Calendar) ([]placeStruct.Calendar, internal.HackError) {
	var bookId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO calendar
(placeId, timeFrom, timeTo, userId) VALUES($1, $2, $3, $4) returning bookId`, body.PlaceId, body.TimeFrom, body.TimeTo, body.UserId).Scan(&bookId)
	if err != nil {
		return []placeStruct.Calendar{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return GetPlaceBookInfo(body.PlaceId)
}

func GetPlaceBookInfo(placeId int64) ([]placeStruct.Calendar, internal.HackError) {
	var result []placeStruct.Calendar
	err := internal.Tools.Connection.Get(&result, `SELECT * FROM calendar WHERE placeId = $1`, placeId)
	if err != nil {
		return []placeStruct.Calendar{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return result, internal.HackError{}
}

func InitTables() internal.HackError {
	_, err := internal.Tools.Connection.Exec(`CREATE TABLE filters (
    	filterId BIGSERIAL PRIMARY KEY NOT NULL ,
    	filterName TEXT NOT NULL
		);`)
	if err != nil {
		log.Println(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	_, err = internal.Tools.Connection.Exec(`CREATE TABLE places (
    placeId BIGSERIAL PRIMARY KEY NOT NULL,
    placeName TEXT NOT NULL,
    filterId BIGINT NOT NULL,
    placeAddress TEXT NOT NULL,
    workingTime TEXT NOT NULL ,
    telephoneNumber TEXT NOT NULL ,
    email TEXT NOT NULL ,
    site TEXT NOT NULL ,
    placeServices TEXT NOT NULL ,
    totalSquare FLOAT NOT NULL , 
    workingSquare FLOAT NOT NULL ,
    commonObjects TEXT NOT NULL , 
    equipment TEXT NOT NULL , 
    rentersCount INTEGER NOT NULL , 
    meta TEXT[] NOT NULL );`)
	if err != nil {
		log.Println(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	_, err = internal.Tools.Connection.Exec(`CREATE TABLE calendar (
    	bookId BIGSERIAL PRIMARY KEY NOT NULL ,
    	placeId BIGINT NOT NULL ,
    	timeFrom DATE NOT NULL,
    	timeTo DATE NOT NULL ,
    	userId BIGINT NOT NULL
		);`)
	if err != nil {
		log.Println(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return internal.HackError{}
}

func DropTable() error {
	_, err := internal.Tools.Connection.Exec(`DROP TABLE filters, places, calendar`)
	if err != nil {
		log.Println(err)
	}
	return err
}
