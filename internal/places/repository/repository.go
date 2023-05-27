package repository

import (
	"database/sql"
	"golang-pkg/internal"
	placeStruct "golang-pkg/internal/places"
	"time"
)

func CreateFilter(body placeStruct.Filter) ([]placeStruct.Filter, *internal.HackError) {
	var filterId int64

	err := internal.Tools.Connection.QueryRowx(`INSERT INTO filters
	(filterName)
	VALUES ($1) returning filterId`, body.FilterName).Scan(&filterId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Filter{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return GetAllFilters()
}

func GetAllFilters() ([]placeStruct.Filter, *internal.HackError) {
	var tmp placeStruct.Filter
	var result []placeStruct.Filter
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM filters`)
	for rows.Next() {
		if err = rows.StructScan(&tmp); err != nil {
			internal.Tools.Logger.Println(err)
			return []placeStruct.Filter{}, &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		result = append(result, tmp)
	}
	internal.Tools.Logger.Println(result)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Filter{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return result, nil
}

func CreatePlace(body placeStruct.Place) (placeStruct.Place, *internal.HackError) {
	var placeId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO places
    (placeName, filterId, placeAddress, workingTime, telephoneNumber, email, site, placeServices, totalSquare, workingSquare,
     commonObjects, equipment, rentersCount, meta) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning placeId`, body.PlaceName,
		body.FilterId, body.PlaceAddress, body.WorkingTime, body.TelephoneNumber, body.Email, body.Site, body.PlaceServices,
		body.TotalSquare, body.WorkingSquare, body.CommonObjects, body.Equipment, body.RentersCount, body.Meta).Scan(&placeId)
	//internal.Tools.Logger.Println(placeId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return placeStruct.Place{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return GetOnePlace(placeId)
}

func GetOnePlace(placeId int64) (placeStruct.Place, *internal.HackError) {
	var body placeStruct.Place
	err := internal.Tools.Connection.Get(&body, `SELECT * FROM places WHERE placeId = $1`, placeId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return placeStruct.Place{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return body, nil
}

func CreateOrder(body placeStruct.Calendar) ([]placeStruct.Calendar, *internal.HackError) {
	var bookId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO calendar
(placeId, timeFrom, timeTo, userId) VALUES($1, $2, $3, $4) returning bookId`, body.PlaceId, body.TimeFrom, body.TimeTo, body.UserId).Scan(&bookId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Calendar{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return GetPlaceBookInfo(body.PlaceId)
}

func GetPlaceBookInfo(placeId int64) ([]placeStruct.Calendar, *internal.HackError) {
	var result []placeStruct.Calendar
	var tmp placeStruct.Calendar
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM calendar WHERE placeId = $1`, placeId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Calendar{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	for rows.Next() {
		if err = rows.StructScan(&tmp); err != nil {
			internal.Tools.Logger.Println(err)
			return []placeStruct.Calendar{}, &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		result = append(result, tmp)
	}

	return result, nil
}

func DeletePlace(placeId int64) *internal.HackError {
	_, err := internal.Tools.Connection.Exec("DELETE FROM places WHERE placeId = $1", placeId)
	if err != nil {
		return &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return nil
}

func UpdatePlace(body placeStruct.Place) *internal.HackError {
	_, err := internal.Tools.Connection.Exec(`UPDATE places SET placeName = $2, filterId = $3, placeAddress = $4, workingTime = $5, 
                  telephoneNumber = $6, email = $7, site = $8, placeServices = $9, 
                  totalSquare = $10, workingSquare = $11, commonObjects = $12, 
                  equipment = $13, rentersCount = $14, meta = $15 WHERE placeId = $1`,
		body.PlaceId, body.PlaceName, body.FilterId, body.PlaceAddress, body.WorkingTime, body.TelephoneNumber,
		body.Email, body.Site, body.PlaceServices, body.TotalSquare, body.WorkingSquare, body.CommonObjects,
		body.Equipment, body.RentersCount, body.Meta)

	if err != nil {
		return &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return nil
}

func SearchPlace(key string) ([]placeStruct.Place, *internal.HackError) {
	var result []placeStruct.Place
	var tmp placeStruct.Place

	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM places WHERE placeName LIKE $1`, "%"+key+"%")
	if err != nil {
		return []placeStruct.Place{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	for rows.Next() {
		if err = rows.StructScan(&tmp); err != nil {
			internal.Tools.Logger.Println(err)
			return []placeStruct.Place{}, &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		result = append(result, tmp)
	}

	return result, nil
}

func DeleteFilter(filterId int64) *internal.HackError {
	_, err := internal.Tools.Connection.Exec("DELETE FROM filters WHERE filterId = $1", filterId)
	if err != nil {
		return &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return nil
}

func CancelOrder(orderId int64, userId int64) *internal.HackError {
	_, err := internal.Tools.Connection.Exec("DELETE FROM calendar WHERE bookId = $1 AND userid = $2", orderId, userId)
	if err != nil {
		return &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return nil
}

func GetPlaces(filterId int, date time.Time, pageNumber int) ([]placeStruct.Place, *internal.HackError) {
	var result []placeStruct.Place
	var body placeStruct.Place
	var placesId []int64
	var tmp int64
	if !date.IsZero() {
		rows, err := internal.Tools.Connection.Queryx(`SELECT placeId FROM calendar WHERE timeFrom < $1 and timeTo > $1`, date)
		if err != nil {
			internal.Tools.Logger.Println(err)
			return []placeStruct.Place{}, &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}

		for rows.Next() {
			if err = rows.Scan(&tmp); err != nil {
				internal.Tools.Logger.Println(err)
			}
			placesId = append(placesId, tmp)
		}

	} else {
		rows, err := internal.Tools.Connection.Queryx(`SELECT placeId FROM places`)
		if err != nil {
			internal.Tools.Logger.Println(err)
			return []placeStruct.Place{}, &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		for rows.Next() {
			if err = rows.Scan(&tmp); err != nil {
				internal.Tools.Logger.Println(err)
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
				internal.Tools.Logger.Println(err)
				return []placeStruct.Place{}, &internal.HackError{
					Code:      500,
					Err:       err,
					Timestamp: time.Now(),
				}
			}
			result = append(result, body)
		} else {
			err := internal.Tools.Connection.Get(&body, `SELECT * FROM places WHERE placeId = $1 and approved = TRUE`, placesId[i])
			if err != nil {
				internal.Tools.Logger.Println(err)
				return []placeStruct.Place{}, &internal.HackError{
					Code:      500,
					Err:       err,
					Timestamp: time.Now(),
				}
			}
			result = append(result, body)
		}
	}
	return result, nil
}

func InitPlaceTables() *internal.HackError {
	_, err := internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  filters (
    	filterId BIGSERIAL PRIMARY KEY NOT NULL ,
    	filterName TEXT NOT NULL
		);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  comments (
    	coomentId BIGSERIAL PRIMARY KEY NOT NULL,
    	placeId BIGSERIAL NOT NULL ,
    	userId BIGSERIAL NOT NULL,
    	comment TEXT NOT NULL,
    	mark FLOAT NOT NULL DEFAULT 0
		);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  places (
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
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  calendar (
    	bookId BIGSERIAL PRIMARY KEY NOT NULL ,
    	placeId BIGINT NOT NULL ,
    	timeFrom DATE NOT NULL,
    	timeTo DATE NOT NULL ,
    	userId BIGINT NOT NULL
		);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  likes (
    	likeId BIGSERIAL PRIMARY KEY NOT NULL ,
    	placeId BIGINT NOT NULL ,
    	userId BIGINT NOT NULL
		);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  users_info (
    user_id       bigserial not null primary key,
    name       text      not null,
    surname    text      not null,
    patronymic text,
    email      text,
    phone text

);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  users_login (
                             login_id bigint references users_info(user_id) primary key,
                             password_hash text
);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  admins (
                        user_id bigint references users_info(user_id) primary key not null,
                        admin_level int
);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	_, err = internal.Tools.Connection.Exec(`CREATE TABLE IF NOT EXISTS  landlords (
                           user_id bigint references users_info(user_id) primary key,
                           post text not null,
                           places bigint[],
                           legal_entity text not null,
                           inn text not null,
                           industry int
);`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	return nil
}

//func DropTable() error {
//	_, err := internal.Tools.Connection.Exec(`DROP TABLE filters, places, calendar, comment, likes`)
//	if err != nil {
//		internal.Tools.Logger.Println(err)
//	}
//	return err
//}

func GetMyOrders(userId int64) ([]placeStruct.Calendar, *internal.HackError) {
	var orders []placeStruct.Calendar
	var tmp placeStruct.Calendar
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM calendar WHERE userid = $1`, userId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Calendar{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	for rows.Next() {
		if err = rows.Scan(&tmp.BookId, &tmp.PlaceId, &tmp.TimeFrom, &tmp.TimeTo, &tmp.UserId); err != nil {
			internal.Tools.Logger.Println(err)
			return []placeStruct.Calendar{}, &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		orders = append(orders, tmp)
	}
	return orders, nil
}

func GetLandPlaces(placesId []int64) ([]placeStruct.LandPlace, *internal.HackError) {
	var result []placeStruct.LandPlace
	var tmp placeStruct.LandPlace
	var tmpPlace placeStruct.Place
	var tmpCalendar placeStruct.Calendar

	for i := 0; i < len(placesId); i++ {
		err := internal.Tools.Connection.Get(&tmpPlace, `SELECT * FROM places WHERE placeId = $1`, placesId[i])
		if err != nil {
			return []placeStruct.LandPlace{}, &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}

		}
		err = internal.Tools.Connection.Get(&tmpCalendar, `SELECT * FROM calendar WHERE placeid = $1`, placesId[i])
		if err != nil {
			return []placeStruct.LandPlace{}, &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}

		}
		tmp.Place = tmpPlace
		tmp.Calendar = tmpCalendar
		result = append(result, tmp)
	}
	return result, nil
}

func GetComments(placeId int64) ([]placeStruct.Comment, *internal.HackError) {
	var result []placeStruct.Comment
	var tmp placeStruct.Comment
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM comments WHERE placeid = $1`, placeId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Comment{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	for rows.Next() {
		if err = rows.Scan(&tmp.CommentId, &tmp.PlaceId, &tmp.UserId, &tmp.Comment, &tmp.Mark); err != nil {
			internal.Tools.Logger.Println(err)
			return []placeStruct.Comment{}, &internal.HackError{
				Code:      500,
				Err:       err,
				Timestamp: time.Now(),
			}
		}
		result = append(result, tmp)
	}
	return result, nil
}

func CreateComment(body placeStruct.Comment) ([]placeStruct.Comment, *internal.HackError) {
	var commentId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO comments
(placeId, userId, comment, mark) VALUES($1, $2, $3, $4) returning commentId`, body.PlaceId, body.UserId, body.Comment, body.Mark).Scan(&commentId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Comment{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	total_mark := 0.0
	err = internal.Tools.Connection.Get(&total_mark, `SELECT AVG(mark) FROM comments WHERE placeid = $1`, body.PlaceId)
	if err != nil {
		return []placeStruct.Comment{}, &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	var mark float64
	err = internal.Tools.Connection.QueryRowx(`UPDATE places SET rating = $1 WHERE placeid = $2 returning mark`, total_mark, body.PlaceId).Scan(&mark)
	if err != nil || mark != total_mark {
		return []placeStruct.Comment{}, &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return GetComments(body.PlaceId)
}

func GetNotApprovedPlaces() ([]placeStruct.Place, *internal.HackError) {
	var result []placeStruct.Place
	var tmp placeStruct.Place
	rows, err := internal.Tools.Connection.Queryx(`SELECT * FROM places WHERE approved = FALSE`)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return []placeStruct.Place{}, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	for rows.Next() {
		if err = rows.Scan(&tmp); err != nil {
			internal.Tools.Logger.Println(err)
		}
		result = append(result, tmp)
	}
	return result, nil
}

func MakeApprove(placeId int64) (placeStruct.Place, *internal.HackError) {
	var approved bool
	err := internal.Tools.Connection.QueryRowx(`UPDATE places SET approved = TRUE WHERE placeid = $1 returning approved`, placeId).Scan(&approved)
	if err != nil {
		return placeStruct.Place{}, &internal.HackError{Code: 500, Err: err, Timestamp: time.Now()}
	}
	return GetOnePlace(placeId)
}

func CreateLike(placeId, userId int64) *internal.HackError {
	var likeId int64
	err := internal.Tools.Connection.QueryRowx(`INSERT INTO likes
    (placeId, userId) VALUES ($1, $2) returning likeId`, placeId, userId).Scan(&likeId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return nil
}

func GetPlaceLikeCount(placeId int64) (int64, *internal.HackError) {
	var likesCount int64
	err := internal.Tools.Connection.Get(&likesCount, `SELECT COUNT(userid) FROM likes WHERE placeid = $1`, placeId)
	if err != nil {
		internal.Tools.Logger.Println(err)
		return 0, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return likesCount, nil
}

func IsLiked(placeId, userId int64) (bool, *internal.HackError) {
	var likesCount int64
	err := internal.Tools.Connection.Get(&likesCount, `SELECT likeid FROM likes WHERE placeid = $1 and userid = $2`, placeId, userId)
	if err != nil || err == sql.ErrNoRows {
		internal.Tools.Logger.Println(err)
		return false, &internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	return likesCount != 0, nil
}
