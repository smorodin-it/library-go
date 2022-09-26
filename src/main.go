package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"library/src/handlers"
	"log"
)

//type DbConfig struct {
//	Host     string
//	UserName string
//	Password string
//	DbName   string
//}

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
		fmt.Sprintf("host=127.0.0.1 user=library password=library dbname=library sslmode=disable"))

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	api := app.Group("/api")

	v1 := api.Group("/v1")

	handlers.ConnectRouter(v1, *db)

	log.Fatal(app.Listen("localhost:3000"))
}
