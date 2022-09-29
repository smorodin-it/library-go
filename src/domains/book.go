package domains

import "time"

type Book struct {
	Id        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Author    string    `json:"author" db:"author"`
	LibraryId string    `json:"libraryId" db:"library_id"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
