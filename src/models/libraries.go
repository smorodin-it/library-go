package models

import (
	"library/src/constants"
	"library/src/forms"
	"library/src/utils"
	"time"
)

type Library struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Address   string    `db:"address"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m *Library) PrepareData(f forms.LibraryAddEditForm, flow int) {
	if flow == constants.ModelCreateFlow {
		m.Id = utils.GenerateUUID()
		m.CreatedAt = time.Now()
	}
	m.Name = f.Name
	m.Address = *f.Address
	m.UpdatedAt = time.Now()
}

type BookInLibrary struct {
	Id          string    `db:"id"`
	LibraryId   string    `db:"library_id"`
	BookId      string    `db:"book_id"`
	AmountTotal int       `db:"amount_total"`
	AmountFact  int       `db:"amount_fact"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (m *BookInLibrary) PrepareData(f *forms.BookToLibraryAddForm, flow int) {
	if flow == constants.ModelCreateFlow {
		m.Id = utils.GenerateUUID()
		m.CreatedAt = time.Now()
	}
	m.LibraryId = f.LibraryId
	m.BookId = f.BookId
	m.AmountTotal = f.AmountTotal
	m.AmountFact = f.AmountFact
	m.UpdatedAt = time.Now()
}
