package user

import (
	"time"

	_ "github.com/go-playground/validator/v10"
)

type User struct {
	Id           string    `db:"id"            json:"id"`
	FirstName    string    `db:"first_name"    json:"first_name"`
	LastName     string    `db:"last_name"     json:"last_name"`
	Email        string    `db:"email"         json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at"    json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"    json:"updated_at"`
}
type RegisterUser struct {
	FirstName    string `db:"first_name"    json:"first_name"    validate:"required"`
	LastName     string `db:"last_name"     json:"last_name"`
	Email        string `db:"email"         json:"email"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
}

type RegisterUserPayload struct {
	FirstName string `db:"first_name" json:"first_name" validate:"required"`
	LastName  string `db:"last_name"  json:"last_name"  validate:"required"`
	Email     string `db:"email"      json:"email"      validate:"required,email"`
	Password  string `db:"password"   json:"password"   validate:"required"`
}

type LoginUser struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"email" json:"password"`
}
