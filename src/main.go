package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type test struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		data := test{
			Message: "world",
		}
		return ctx.Status(http.StatusOK).JSON(data)
	})

	log.Fatal(app.Listen("localhost:3000"))
}
