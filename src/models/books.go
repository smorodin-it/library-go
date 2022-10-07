package models

import "time"

type Book struct {
	Id        string    `db:"id"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	LibraryId string    `db:"library_id"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
