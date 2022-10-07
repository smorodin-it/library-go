package models

import "time"

type User struct {
	Id           string    `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	ProfileId    string    `db:"profile_id"`
	Active       bool      `db:"active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Profile struct {
	Id         string    `db:"id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Patronymic string    `db:"patronymic"`
	Phone      string    `db:"phone"`
	Address    string    `db:"phone"`
	UserId     string    `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
