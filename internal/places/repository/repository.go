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
		log.Println(err)
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
		log.Println(err)
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
	log.Println(placeId)
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
		log.Println(err)
		return placeStruct.Place{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return body, internal.HackError{}
}

func CreateOrder(body placeStruct.Calendar) ([]placeStruct.Calendar, internal.HackError) {
	var bookId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO calendar
(placeId, timeFrom, timeTo, userId) VALUES($1, $2, $3, $4) returning bookId`, body.PlaceId, body.TimeFrom, body.TimeTo, body.UserId).Scan(&bookId)
	if err != nil {
		log.Println(err)
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
	var tmp placeStruct.Calendar
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM calendar WHERE placeId = $1`, placeId)
	if err != nil {
		log.Println(err)
		return []placeStruct.Calendar{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	for rows.Next() {
		if err = rows.StructScan(&tmp); err != nil {
			log.Println(err)
			return []placeStruct.Calendar{}, internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		result = append(result, tmp)
	}

	return result, internal.HackError{}
}

func DeletePlace(placeId int64) internal.HackError {
	_, err := internal.Tools.Connection.Exec("DELETE FROM places WHERE placeId = $1", placeId)
	if err != nil {
		return internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return internal.HackError{}
}

func UpdatePlace(body placeStruct.Place) internal.HackError {
	_, err := internal.Tools.Connection.Exec(`UPDATE places SET placeName = $2, filterId = $3, placeAddress = $4, workingTime = $5, 
                  telephoneNumber = $6, email = $7, site = $8, placeServices = $9, 
                  totalSquare = $10, workingSquare = $11, commonObjects = $12, 
                  equipment = $13, rentersCount = $14, meta = $15 WHERE placeId = $1`,
		body.PlaceId, body.PlaceName, body.FilterId, body.PlaceAddress, body.WorkingTime, body.TelephoneNumber,
		body.Email, body.Site, body.PlaceServices, body.TotalSquare, body.WorkingSquare, body.CommonObjects,
		body.Equipment, body.RentersCount, body.Meta)

	if err != nil {
		return internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return internal.HackError{}
}

func SearchPlace(key string) ([]placeStruct.Place, internal.HackError) {
	var result []placeStruct.Place
	var tmp placeStruct.Place

	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM places WHERE placeName LIKE $1`, "%"+key+"%")
	if err != nil {
		return []placeStruct.Place{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	for rows.Next() {
		if err = rows.StructScan(&tmp); err != nil {
			log.Println(err)
			return []placeStruct.Place{}, internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		result = append(result, tmp)
	}

	return result, internal.HackError{}
}

func DeleteFilter(filterId int64) internal.HackError {
	_, err := internal.Tools.Connection.Exec("DELETE FROM filters WHERE filterId = $1", filterId)
	if err != nil {
		return internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return internal.HackError{}
}

func CancelOrder(orderId int64, userId int64) internal.HackError {
	_, err := internal.Tools.Connection.Exec("DELETE FROM calendar WHERE bookId = $1 AND userid = $2", orderId, userId)
	if err != nil {
		return internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return internal.HackError{}
}

func GetPlaces(filterId int, date time.Time, pageNumber int) ([]placeStruct.Place, internal.HackError) {
	var result []placeStruct.Place
	var body placeStruct.Place
	var placesId []int64
	var tmp int64
	if !date.IsZero() {
		rows, err := internal.Tools.Connection.Queryx(`SELECT placeId FROM calendar WHERE timeFrom < $1 and timeTo > $1`, date)
		if err != nil {
			log.Println(err)
			return []placeStruct.Place{}, internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}

		for rows.Next() {
			if err = rows.Scan(&tmp); err != nil {
				log.Println(err)
			}
			placesId = append(placesId, tmp)
		}

	} else {
		rows, err := internal.Tools.Connection.Queryx(`SELECT placeId FROM places`)
		if err != nil {
			log.Println(err)
			return []placeStruct.Place{}, internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		for rows.Next() {
			if err = rows.Scan(&tmp); err != nil {
				log.Println(err)
			}
			placesId = append(placesId, tmp)
		}

	}

	lim := 0
	if pageNumber*10 > len(placesId) {
		lim = len(placesId)
	} else {
		lim = pageNumber
	}

	for i := pageNumber*10 - 10; i < lim; i++ {
		if filterId != 0 {
			err := internal.Tools.Connection.Get(&body, `SELECT * FROM places WHERE placeId = $1 and filterId = $2 and approved = TRUE`, placesId[i], filterId)
			if err != nil {
				log.Println(err)
				return []placeStruct.Place{}, internal.HackError{
					Code:      500,
					Err:       err,
					Timestamp: time.Now(),
				}
			}
			result = append(result, body)
		} else {
			err := internal.Tools.Connection.Get(&body, `SELECT * FROM places WHERE placeId = $1 and approved = TRUE`, placesId[i])
			if err != nil {
				log.Println(err)
				return []placeStruct.Place{}, internal.HackError{
					Code:      500,
					Err:       err,
					Timestamp: time.Now(),
				}
			}
			result = append(result, body)
		}
	}
	return result, internal.HackError{}
}

func InitPlaceTables() internal.HackError {
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
	_, err = internal.Tools.Connection.Exec(`CREATE TABLE comments (
    	coomentId BIGSERIAL PRIMARY KEY NOT NULL,
    	placeId BIGSERIAL NOT NULL ,
    	userId BIGSERIAL NOT NULL,
    	comment TEXT NOT NULL,
    	mark FLOAT NOT NULL DEFAULT 0
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
    meta TEXT[] NOT NULL,
    rating FLOAT NOT NULL DEFAULT 0,
    approved BOOL DEFAULT FALSE
     );`)
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

func GetMyOrders(userId int64) ([]placeStruct.Calendar, internal.HackError) {
	var orders []placeStruct.Calendar
	var tmp placeStruct.Calendar
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM calendar WHERE userid = $1`, userId)
	if err != nil {
		log.Println(err)
		return []placeStruct.Calendar{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	for rows.Next() {
		if err = rows.Scan(&tmp.BookId, &tmp.PlaceId, &tmp.TimeFrom, &tmp.TimeTo, &tmp.UserId); err != nil {
			log.Println(err)
			return []placeStruct.Calendar{}, internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		orders = append(orders, tmp)
	}
	return orders, internal.HackError{}
}

func GetLandPlaces(placesId []int64) ([]placeStruct.LandPlace, internal.HackError) {
	var result []placeStruct.LandPlace
	var tmp placeStruct.LandPlace
	var tmpPlace placeStruct.Place
	var tmpCalendar placeStruct.Calendar

	for i := 0; i < len(placesId); i++ {
		err := internal.Tools.Connection.Get(&tmpPlace, `SELECT * FROM places WHERE placeId = $1`, placesId[i])
		if err != nil {
			return []placeStruct.LandPlace{}, internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}

		}
		err = internal.Tools.Connection.Get(&tmpCalendar, `SELECT * FROM calendar WHERE placeid = $1`, placesId[i])
		if err != nil {
			return []placeStruct.LandPlace{}, internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}

		}
		tmp.Place = tmpPlace
		tmp.Calendar = tmpCalendar
		result = append(result, tmp)
	}
	return result, internal.HackError{}
}

func GetComments(placeId int64) ([]placeStruct.Comment, internal.HackError) {
	var result []placeStruct.Comment
	var tmp placeStruct.Comment
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM comments WHERE placeid = $1`, placeId)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	for rows.Next() {
		if err = rows.Scan(&tmp.CommentId, &tmp.PlaceId, &tmp.UserId, &tmp.Comment, &tmp.Mark); err != nil {
			log.Println(err)
			return []placeStruct.Comment{}, internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		result = append(result, tmp)
	}
	return result, internal.HackError{}
}

func CreateComment(body placeStruct.Comment) ([]placeStruct.Comment, internal.HackError) {
	var commentId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO comments
(placeId, userId, comment, mark) VALUES($1, $2, $3, $4) returning commentId`, body.PlaceId, body.UserId, body.Comment, body.Mark).Scan(&commentId)
	if err != nil {
		log.Println(err)
		return []placeStruct.Comment{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	total_mark := 0.0
	err = internal.Tools.Connection.Get(&total_mark, `SELECT AVG(mark) FROM comments WHERE placeid = $1`, body.PlaceId)
	if err != nil {
		return []placeStruct.Comment{}, internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	var mark float64
	err = internal.Tools.Connection.QueryRowx(`UPDATE places SET rating = $1 WHERE placeid = $2 returning mark`, total_mark, body.PlaceId).Scan(&mark)
	if err != nil || mark != total_mark {
		return []placeStruct.Comment{}, internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return GetComments(body.PlaceId)
}

func GetNotApprovedPlaces() ([]placeStruct.Place, internal.HackError) {
	var result []placeStruct.Place
	var tmp placeStruct.Place
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM places WHERE approved = FALSE`)
	if err != nil {
		log.Println(err)
		return []placeStruct.Place{}, internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	for rows.Next() {
		if err = rows.Scan(&tmp); err != nil {
			log.Println(err)
		}
		result = append(result, tmp)
	}
	return result, internal.HackError{}
}

func MakeApprove(placeId int64) (placeStruct.Place, internal.HackError) {
	var approved bool
	err := internal.Tools.Connection.QueryRowx(`UPDATE places SET approved = TRUE WHERE placeid = $1 returning approved`, placeId).Scan(&approved)
	if err != nil {
		return placeStruct.Place{}, internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return GetOnePlace(placeId)
}
