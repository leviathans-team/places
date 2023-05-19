package repository

//
//import (
//	models "golang-pkg/internal"
//	"log"
//)
//
//func CreateFilter(body *models.Filter) {
//	models.Connection.Database.QueryRowx(`INSERT INTO filters
//	(filterName)
//	VALUES ($1)`, body.FilterName)
//}
//
//func GetAllFilters() []models.Filter {
//	var result []models.Filter
//	models.Connection.Database.Get(&result, `SELECT * FROM filters`)
//	return result
//}

//
//func CreateType(types string) {
//	marketplace.Connection.Database.QueryRowx(`INSERT INTO types
//	(producttypes)
//	VALUES ($1)`, types)
//}
//
//func GetItem(id int) marketplace.Items {
//	var body marketplace.Items
//	err := marketplace.Connection.Database.Get(&body, `SELECT * FROM items WHERE item_id = $1`, id)
//	if err != nil {
//		log.Println(err)
//	}
//	return body
//}
//
//func GetItems(start int, types int) []marketplace.Items {
//	body := make([]marketplace.Items, 0)
//	var rows *sqlx.Rows
//	var err error
//	var stru marketplace.Items
//	if types == 0 {
//		rows, err = marketplace.Connection.Database.Queryx(`SELECT * FROM items`)
//	} else {
//		rows, err = marketplace.Connection.Database.Queryx(`SELECT * FROM items WHERE product_type = $1`, types)
//	}
//	if err != nil {
//		log.Println(err)
//	}
//	for rows.Next() {
//
//		if err := rows.Scan(&stru.Item_id, &stru.Name, &stru.Description, &stru.Photo_id, &stru.Characteristics,
//			&stru.Price, &stru.Count, &stru.Product_type); err != nil {
//			log.Println(err)
//		}
//		body = append(body, stru)
//	}
//	return body
//
//}
//
//func Del(id int) error {
//	var err error
//	_, err = marketplace.Connection.Database.Exec("DELETE FROM items WHERE item_id = $1", id)
//	return err
//}
//
//func GetTypes() []marketplace.ProductType {
//	types := make([]marketplace.ProductType, 0)
//	var stru marketplace.ProductType
//	rows, err := marketplace.Connection.Database.Queryx(`SELECT * FROM types`)
//	if err != nil {
//		log.Println(err)
//	}
//	for rows.Next() {
//
//		if err := rows.Scan(&stru.Type_id, &stru.Product_types); err != nil {
//			log.Println(err)
//		}
//		types = append(types, stru)
//	}
//	//err := marketplace.Connection.Database.Select(&types, `SELECT * FROM types`) // - то же самое что сверху только не работает
//	//if err != nil {
//	//	log.Println(err)
//	//}
//	return types
//}
//
//func UpdateById(body marketplace.Items) marketplace.Items {
//	var new marketplace.Items
//
//	_, err := marketplace.Connection.Database.Queryx(`UPDATE items SET name = $1, description = $2, photo_id = $3,
//                 characteristics = $4, price = $5, count = $6, product_type = $7 WHERE item_id = $8`, body.Name,
//		body.Description, body.Photo_id, body.Characteristics, body.Price, body.Count, body.Product_type,
//		body.Item_id)
//	if err != nil {
//		log.Println(err)
//	}
//	marketplace.Connection.Database.Get(&new, `SELECT * FROM items WHERE item_id = $1`, body.Item_id)
//	return new
//}
//
//func InitTables() error {
//	_, err := models.Connection.Database.Exec(`CREATE TABLE filters (
//    	filterId BIGSERIAL PRIMARY KEY NOT NULL ,
//    	filterName TEXT NOT NULL ,
//		);`)
//	if err != nil {
//		log.Println(err)
//	}
//	return err
//}
//
//func DropTable() error {
//	_, err := models.Connection.Database.Exec(`DROP TABLE filters`)
//	if err != nil {
//		log.Println(err)
//	}
//	return err
//}
