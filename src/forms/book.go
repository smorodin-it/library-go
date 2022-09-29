package forms

import (
	"errors"
	"unicode/utf8"
)

type BookAddEditForm struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (f *BookAddEditForm) Validate() error {
	if &f.Title == nil || f.Title == "" {
		return errors.New("title field is required")
	}
	if utf8.RuneCountInString(f.Title) > 255 {
		return errors.New("field max length 255 chars")
	}
	if &f.Author == nil || f.Author == "" {
		return errors.New("author field is required")
	}
	if utf8.RuneCountInString(f.Author) > 255 {
		return errors.New("field max length 255 chars")
	}
	return nil
}
