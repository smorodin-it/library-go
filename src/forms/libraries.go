package forms

import (
	"errors"
	"github.com/go-playground/validator"
	"library/src/responses"
	"unicode/utf8"
)

type LibraryAddEditForm struct {
	Name    string  `json:"name"`
	Address *string `json:"address"`
}

func (f *LibraryAddEditForm) Validate() error {
	if &f.Name == nil || f.Name == "" {
		return errors.New("name field is required")
	}
	if utf8.RuneCountInString(f.Name) > 255 {
		return errors.New("field max length 255 chars")
	}
	if utf8.RuneCountInString(*f.Address) > 2000 {
		return errors.New("field max length 2000 chars")
	}
	return nil
}

type BookToLibraryAddForm struct {
	LibraryId   string `json:"libraryId" validate:"required,uuid4"`
	BookId      string `json:"bookId" validate:"required,uuid4"`
	AmountTotal int    `json:"amountTotal" validate:"required,numeric"`
	AmountFact  int    `json:"amountFact" validate:"omitempty,numeric"`
}

func (f *BookToLibraryAddForm) Validate() []*responses.ErrorResponse {
	validate := validator.New()
	var errs []*responses.ErrorResponse
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element responses.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errs = append(errs, &element)
		}
		return errs
	}
	return nil
}
