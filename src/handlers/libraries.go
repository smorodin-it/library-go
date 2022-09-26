package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"library/src/constants"
	"library/src/domains"
	"library/src/utils"
	"log"
	"net/http"
)

func ListLibraries(ctx *fiber.Ctx, db *sqlx.DB) error {
	var libraries []domains.Library
	sql := fmt.Sprintf("SELECT * FROM %s", constants.LibraryTable)

	err := db.Select(&libraries, sql)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(libraries)
}

func RetrieveLibrary(ctx *fiber.Ctx, db *sqlx.DB) error {
	id := ctx.Params("libraryId")
	library := domains.Library{}

	sql := fmt.Sprintf("SELECT * from %s WHERE id=$1", constants.LibraryTable)
	err := db.Get(&library, sql, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(library)

}

func CreateLibrary(ctx *fiber.Ctx, db *sqlx.DB) error {
	model := new(domains.Library)

	if err := ctx.BodyParser(model); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	model.Id = utils.GenerateUUID()

	sql := fmt.Sprintf("INSERT INTO %s (id, name, address) VALUES ($1, $2, $3)", constants.LibraryTable)

	res, err := db.Query(sql, model.Id, model.Name, model.Address)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	log.Println(res)

	return ctx.Status(http.StatusCreated).JSON(model)
}
