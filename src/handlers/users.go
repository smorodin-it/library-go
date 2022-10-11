package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"library/src/constants"
	"library/src/forms"
	"library/src/models"
	"net/http"
)

func Login(ctx *fiber.Ctx, db *sqlx.DB) error {
	form := new(forms.UserSignInForm)
	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	//	TODO: Form validation

	user := new(models.User)

	sql := fmt.Sprintf("SELECT * from %s WHERE email=$1", constants.LibraryTable)

	err := db.Get(&user, sql, form.Email)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// TODO: Return tokens
	return ctx.SendStatus(http.StatusOK)
}

func RefreshToken(ctx *fiber.Ctx, db *sqlx.DB) error {
	//	TODO: Implement
	return ctx.Status(http.StatusInternalServerError).SendString("Method not implemented!")
}

func Logout(ctx *fiber.Ctx, db *sqlx.DB) error {
	//	TODO: Implement
	return ctx.Status(http.StatusInternalServerError).SendString("Method not implemented!")
}
