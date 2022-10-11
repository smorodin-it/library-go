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

func (m *Library) GetData(form forms.LibraryAddEditForm, modelFlow int) Library {
	if modelFlow == constants.ModelCreateFlow {
		m.Id = utils.GenerateUUID()
	}

	m.Name = form.Name
	m.Address = *form.Address
	m.UpdatedAt = time.Now()

	return *m
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
