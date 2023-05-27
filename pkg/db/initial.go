package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang-pkg/config"
)

func InitPsqlDB(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		*c.Postgres.Host,
		*c.Postgres.Port,
		*c.Postgres.User,
		*c.Postgres.Password,
		*c.Postgres.DbName)
	fmt.Println(connectionUrl)
	database, err := sqlx.Connect("pgx", connectionUrl)
	if err != nil {
		return nil, err
	}
	return database, nil
}
