package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-pkg/config"
	models "golang-pkg/internal"
	"golang-pkg/internal/delivery"
	"golang-pkg/pkg/db"
	"log"
)

func main() {
	var app = fiber.New()
	viperConf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}
	models.Connection.Database, err = db.InitPsqlDB(conf)
	if err != nil {
		log.Fatal(err)
	}
	// открываем соединение mongo
	//session, err := mgo.Dial("mongodb://127.0.0.1")
	//if err != nil {
	//	panic(err)
	//}
	//defer session.Close()
	//
	//// получаем коллекцию
	//productCollection := session.DB("productdb").C("products")
	//// критерий выборки
	//query := bson.M{}
	//// объект для сохранения результата
	//products := []Product{}
	//productCollection.Find(query).All(&products)

	//middleware.InitTables()
	delivery.Hearing(app)
}
