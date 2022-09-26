package forms

import (
	"errors"
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
