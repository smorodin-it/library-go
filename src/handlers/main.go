package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"library/src/constants"
)

func ConnectRouter(app fiber.Router, db sqlx.DB) {
	// Libraries routes
	app.Get("/libraries", func(ctx *fiber.Ctx) error {
		return ListLibraries(ctx, &db)
	})

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

	app.Post("/libraries/batch", func(ctx *fiber.Ctx) error {
		return CreateLibrariesInBatch(ctx, &db)
	})

	// Books routes
	app.Get("/books", func(ctx *fiber.Ctx) error {
		return ListBooks(ctx, &db)
	})

	app.Get(fmt.Sprintf("/book/:%s", constants.BookIdField), func(ctx *fiber.Ctx) error {
		return RetrieveBook(ctx, &db)
	})

	app.Post("/book", func(ctx *fiber.Ctx) error {
		return CreateBook(ctx, &db)
	})

	app.Put(fmt.Sprintf("/book/:%s", constants.BookIdField), func(ctx *fiber.Ctx) error {
		return UpdateBook(ctx, &db)
	})

	app.Patch(fmt.Sprintf("/book/activate/:%s", constants.BookIdField), func(ctx *fiber.Ctx) error {
		return ToggleBookActive(ctx, &db, true)
	})

	app.Patch(fmt.Sprintf("/book/deactivate/:%s", constants.BookIdField), func(ctx *fiber.Ctx) error {
		return ToggleBookActive(ctx, &db, false)
	})
}
