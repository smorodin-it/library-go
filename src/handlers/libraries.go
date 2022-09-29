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
	"strconv"
	"time"
)

func ListLibraries(ctx *fiber.Ctx, db *sqlx.DB) error {
	var libraries []domains.Library

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(ctx.Query("perPage"))
	if err != nil {
		perPage = constants.PerPageDefault
	}

	sql := fmt.Sprintf("select * from %s order by name asc limit $1 offset $2", constants.LibraryTable)

	println(page, perPage)

	err = db.Select(&libraries, sql, perPage, (page-1)*perPage)
	//err = db.Select(&libraries, sql, 20, 0)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(libraries)
}

func RetrieveLibrary(ctx *fiber.Ctx, db *sqlx.DB) error {
	id := ctx.Params(constants.LibraryIdField)
	library := domains.Library{}

	sql := fmt.Sprintf("SELECT * from %s WHERE id=$1", constants.LibraryTable)
	err := db.Get(&library, sql, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(library)

}

func CreateLibrary(ctx *fiber.Ctx, db *sqlx.DB) error {
	form := new(forms.LibraryAddEditForm)

	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := form.Validate(); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	model := new(domains.Library)

	model.Id = utils.GenerateUUID()
	model.Name = form.Name
	model.Address = *form.Address
	model.UpdatedAt = time.Now()

	sql := fmt.Sprintf("insert into %s (id, name, address) values ($1, $2, $3)", constants.LibraryTable)

	_, err := db.Query(sql, model.Id, model.Name, model.Address)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(struct {
		Id string `json:"id"`
	}{
		Id: model.Id,
	})
}

func UpdateLibrary(ctx *fiber.Ctx, db *sqlx.DB) error {
	id := ctx.Params(constants.LibraryIdField)

	form := new(forms.LibraryAddEditForm)

	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := form.Validate(); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	model := new(domains.Library)

	model.Name = form.Name
	model.Address = *form.Address
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

func ToggleLibraryActive(ctx *fiber.Ctx, db *sqlx.DB, active bool) error {
	id := ctx.Params(constants.LibraryIdField)

	sql := fmt.Sprintf("update %s set active=$1, updated_at=$2 where id=$3", constants.LibraryTable)
	_, err := db.Query(sql, active, time.Now(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(struct {
		Status bool `json:"status"`
	}{
		Status: true,
	})
}

func CreateLibrariesInBatch(ctx *fiber.Ctx, db *sqlx.DB) error {
	var libraries []domains.Library

	if err := ctx.BodyParser(&libraries); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	for _, l := range libraries {
		form := forms.LibraryAddEditForm{
			Name:    l.Name,
			Address: &l.Address,
		}
		if err := form.Validate(); err != nil {
			return ctx.Status(http.StatusBadRequest).SendString(err.Error())
		}

		model := new(domains.Library)

		model.Id = utils.GenerateUUID()
		model.Name = form.Name
		model.Address = *form.Address

		sql := fmt.Sprintf("insert into %s (id, name, address) values ($1, $2, $3)", constants.LibraryTable)
		_, err := db.Query(sql, model.Id, model.Name, model.Address)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}
	}

	return ctx.Status(http.StatusOK).JSON(struct {
		Status bool `json:"status"`
	}{
		Status: true,
	})
}
