package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"library/src/constants"
	"library/src/forms"
	"library/src/models"
	"library/src/utils"
	"net/http"
	"strconv"
	"time"
)

func ListBooks(ctx *fiber.Ctx, db *sqlx.DB) error {
	var books []models.Book

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(ctx.Query("perPage"))
	if err != nil {
		perPage = constants.PerPageDefault
	}

	sql := fmt.Sprintf("select * from %s order by title asc limit $1 offset $2", constants.BookTable)

	err = db.Select(&books, sql, perPage, (page-1)*perPage)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(books)
}

func RetrieveBook(ctx *fiber.Ctx, db *sqlx.DB) error {
	id := ctx.Params(constants.BookIdField)
	library := models.Book{}

	sql := fmt.Sprintf("SELECT * from %s WHERE id=$1", constants.BookTable)
	err := db.Get(&library, sql, id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(library)

}

func CreateBook(ctx *fiber.Ctx, db *sqlx.DB) error {
	form := new(forms.BookAddEditForm)

	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := form.Validate(); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	model := new(models.Book)

	model.Id = utils.GenerateUUID()
	model.Title = form.Title
	model.Author = form.Author
	model.UpdatedAt = time.Now()

	sql := fmt.Sprintf("insert into %s (id, title, author) values ($1, $2, $3)", constants.BookTable)

	_, err := db.Query(sql, model.Id, model.Title, model.Author)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(struct {
		Id string `json:"id"`
	}{
		Id: model.Id,
	})
}

func UpdateBook(ctx *fiber.Ctx, db *sqlx.DB) error {
	id := ctx.Params(constants.BookIdField)

	form := new(forms.BookAddEditForm)

	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := form.Validate(); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	model := new(models.Book)

	model.Title = form.Title
	model.Author = form.Author
	model.UpdatedAt = time.Now()

	sql := fmt.Sprintf("update %s set title=$1, author=$2, updated_at=$3 where id=$4", constants.BookTable)

	_, err := db.Query(sql, model.Title, model.Author, model.UpdatedAt, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(struct {
		Status bool `json:"status"`
	}{
		Status: true,
	})
}

func ToggleBookActive(ctx *fiber.Ctx, db *sqlx.DB, active bool) error {
	id := ctx.Params(constants.BookIdField)

	sql := fmt.Sprintf("update %s set active=$1, updated_at=$2 where id=$3", constants.BookTable)
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
