package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	//"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"library/src/constants"
)

func ConnectRouter(app fiber.Router, db sqlx.DB) {

	/*
		Auth Routes
	*/

	/*
		Public Routes
	*/

	// Libraries routes
	app.Get("/libraries", func(ctx *fiber.Ctx) error {
		return ListLibrariesAdmin(ctx, &db)
	})

	app.Get(fmt.Sprintf("/library/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return RetrieveLibraryAdmin(ctx, &db)
	})

	app.Post("/library", func(ctx *fiber.Ctx) error {
		return CreateLibraryAdmin(ctx, &db)
	})

	app.Put(fmt.Sprintf("/library/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return UpdateLibraryAdmin(ctx, &db)
	})

	app.Patch(fmt.Sprintf("/library/activate/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return ToggleLibraryActiveAdmin(ctx, &db, true)
	})

	app.Patch(fmt.Sprintf("/library/deactivate/:%s", constants.LibraryIdField), func(ctx *fiber.Ctx) error {
		return ToggleLibraryActiveAdmin(ctx, &db, false)
	})

	app.Post("/libraries/batch", func(ctx *fiber.Ctx) error {
		return CreateLibrariesInBatchAdmin(ctx, &db)
	})

	app.Post("/library/add-book", func(ctx *fiber.Ctx) error {
		return AddBookToLibraryAdmin(ctx, &db)
	})

	app.Get("/library/:libraryId/list-books", func(ctx *fiber.Ctx) error {
		return ListAllBooksInLibraryAdmin(ctx, &db)
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

	/*
		Private Routes
	*/
	app.Use(jwtware.New(jwtware.Config{
		// TODO: Move to env
		SigningKey: []byte("SOME_VERY_SECRET_KEY"),
	}))
}
