package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html"
	"golang-pkg/config"
	_ "golang-pkg/docs"
	models "golang-pkg/internal"
	"golang-pkg/internal/auth/handlers"
	"golang-pkg/internal/places/delivery"
	"golang-pkg/internal/places/repository"
	userHandlers "golang-pkg/internal/user/handlers"
	"golang-pkg/pkg/db"
	"log"
	"os"
)

// @title Hack
// @version 1.0
// @description Документация API
// @host 37.18.110.184:3000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// swag init -g cmd/main.go --pd
func main() {
	viperConf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}

	//models.Tools.Logger = logger.NewServiceLogger(conf)
	models.Tools.Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
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

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	handlers.SetupRoutesForAuth(app)
	userHandlers.UserPanel(app)
	delivery.Hearing(app)

	repository.InitPlaceTables()
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
