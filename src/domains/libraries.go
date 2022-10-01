package domains

import "time"

type Library struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Address   string    `json:"address" db:"address"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type BookInLibrary struct {
	Id          string    `json:"id" db:"id"`
	LibraryId   string    `json:"libraryId" db:"library_id"`
	BookId      string    `json:"bookId" db:"book_id"`
	AmountTotal int       `json:"amountTotal" db:"amount_total"`
	AmountFact  int       `json:"amountFact" db:"amount_fact"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
