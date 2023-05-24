package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"golang-pkg/config"
	models "golang-pkg/internal"
	"golang-pkg/internal/auth/handlers"
	"golang-pkg/internal/places/delivery"
	"golang-pkg/internal/places/repository"
	"golang-pkg/pkg/db"
	"golang-pkg/pkg/logger"
	"log"
	"os"
)

func main() {
	viperConf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}

	models.Tools.Logger = logger.NewServiceLogger(conf)
	models.Tools.AdminLogger = log.New(os.Stdout, "ADMINLOG: ", log.LstdFlags)

	models.Tools.Connection, err = db.InitPsqlDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = models.Tools.Connection.Ping()
	if err != nil {
		log.Panic(err)
	}
	//кваврп

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	handlers.SetupRoutesForAuth(app)
	delivery.Hearing(app) // создай группу для сових ручек, в будующем будет проще поддерживать/фиксить/строить код

	repository.InitPlaceTables()
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
