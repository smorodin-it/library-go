package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"library/src/constants"
	"library/src/domains"
	"library/src/forms"
	"library/src/utils"
	"net/http"
	"time"
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

	_, err := db.Query(sql, model.Id, model.Name, model.Address)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(struct {
		Id string `json:"id"`
	}{
		Id: model.Id,
	})
}

func UpdateLibrary(ctx *fiber.Ctx, db *sqlx.DB) error {
	id := ctx.Params("libraryId")

	form := new(forms.LibraryAddEditForm)

	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := form.Validate(); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	model := new(domains.Library)

	model.Name = form.Name
	model.Address = form.Address
	model.UpdatedAt = time.Now()

	sql := fmt.Sprintf("update %s set name=$1, address=$2, updated_at=$3 where id=$4", constants.LibraryTable)

	_, err := db.Query(sql, model.Name, model.Address, model.UpdatedAt, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(struct {
		Status bool `json:"status"`
	}{
		Status: true,
	})
}
