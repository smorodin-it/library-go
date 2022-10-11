package models

import (
	"golang.org/x/crypto/bcrypt"
	"library/src/constants"
	"library/src/forms"
	"library/src/utils"
	"time"
)

type User struct {
	Id           string    `db:"id"`
	Email        string    `db:"email"`
	PasswordHash []byte    `db:"password_hash"`
	ProfileId    string    `db:"profile_id"`
	Active       bool      `db:"active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (m *User) PrepareData(f *forms.UserSignUpForm, flow int) {
	if flow == constants.ModelCreateFlow {
		m.Id = utils.GenerateUUID()
		m.CreatedAt = time.Now()
	}
	m.Email = f.Email

	m.UpdatedAt = time.Now()
}

func (m *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.PasswordHash = hash
	return nil
}

func (m *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(m.PasswordHash, []byte(password))
	return err != nil
}

type Profile struct {
	Id         string    `db:"id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Patronymic *string   `db:"patronymic"`
	Phone      string    `db:"phone"`
	Address    string    `db:"phone"`
	UserId     string    `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (m *Profile) PrepareData(f *forms.ProfileAddEditForm, flow int) {
	if flow == constants.ModelCreateFlow {
		m.Id = utils.GenerateUUID()
		m.CreatedAt = time.Now()
	}
	m.FirstName = f.FirstName
	m.LastName = f.LastName
	m.Patronymic = f.Patronymic
	m.Phone = f.Phone
	m.Address = f.Address
	m.UserId = f.UserId
	m.UpdatedAt = time.Now()
}
