package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func ConnectRouter(app fiber.Router, db sqlx.DB) {
	// Libraries routes
	// List
	app.Get("/libraries", func(ctx *fiber.Ctx) error {
		return ListLibraries(ctx, &db)
	})

	// Retrieve
	app.Get("/library/:libraryId", func(ctx *fiber.Ctx) error {
		return RetrieveLibrary(ctx, &db)
	})

	app.Post("/library", func(ctx *fiber.Ctx) error {
		return CreateLibrary(ctx, &db)
	})
}
