package models

import (
	"library/src/constants"
	"library/src/forms"
	"library/src/utils"
	"time"
)

type Book struct {
	Id     string `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	//LibraryId string    `db:"library_id"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

func (m *Book) PrepareData(f *forms.BookAddEditForm, flow int) {
	if flow == constants.ModelCreateFlow {
		m.Id = utils.GenerateUUID()
		m.CreatedAt = time.Now()
	}
	m.Title = f.Title
	m.Author = f.Author
	m.UpdatedAt = time.Now()
}
