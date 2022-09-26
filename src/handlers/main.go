package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"library/src/constants"
)

func ConnectRouter(app fiber.Router, db sqlx.DB) {
	// Libraries routes
	// List
	app.Get("/libraries", func(ctx *fiber.Ctx) error {
		return ListLibraries(ctx, &db)
	})

	// Retrieve
	app.Get(fmt.Sprintf("/library/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return RetrieveLibrary(ctx, &db)
	})

	app.Post("/library", func(ctx *fiber.Ctx) error {
		return CreateLibrary(ctx, &db)
	})

	app.Put(fmt.Sprintf("/library/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return UpdateLibrary(ctx, &db)
	})

	app.Patch(fmt.Sprintf("/library/activate/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return ToggleLibraryActive(ctx, &db, true)
	})

	app.Patch(fmt.Sprintf("/library/deactivate/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return ToggleLibraryActive(ctx, &db, false)
	})
}
