package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"library/src/constants"
	"library/src/domains"
	"log"
	"net/http"
)

type DbConfig struct {
	Host     string
	UserName string
	Password string
	DbName   string
}

func main() {
	//err := godotenv.Load()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//dbConfig := DbConfig{
	//	Host:     os.Getenv("DB_HOST"),
	//	UserName: os.Getenv("DB_USERNAME"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//	DbName:   os.Getenv("DB_NAME"),
	//}

	db, err := sqlx.Connect("postgres",
		//fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbConfig.Host, dbConfig.UserName, dbConfig.Password, dbConfig.DbName))
		fmt.Sprintf("host=localhost user=library password=library dbname=library sslmode=disable"))

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get(
		"/libraries",
		func(ctx *fiber.Ctx) error {
			var libraries []domains.Library
			err := db.Select(&libraries, fmt.Sprintf("SELECT * FROM %s", constants.LibraryTable))
			if err != nil {
				return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
			}
			return ctx.Status(http.StatusOK).JSON(libraries)
		},
	)

	log.Fatal(app.Listen("localhost:3000"))
}
