package models

import (
	"regexp"
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
	FirstName    string `db:"first_name"    json:"first_name"`
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

func (r *RegisterUserPayload) Validate() (errorMap map[string][]string) {
	errorMap = make(map[string][]string)
	if r.Email == "" {
		errorMap["email"] = append(errorMap["email"], "Email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(r.Email) {
		errorMap["email"] = append(errorMap["email"], "Invalid email format")
	}
	if r.Password == "" {
		errorMap["password"] = append(errorMap["password"], "Password is required")
	}
	if len(r.Password) < 6 {
		errorMap["password"] = append(
			errorMap["password"],
			"Password must be at least 6 characters",
		)
	}
	if r.FirstName == "" {
		errorMap["first_name"] = append(errorMap["first_name"], "First name is required")
	}
	if r.FirstName != "" && len(r.FirstName) < 2 {
		errorMap["first_name"] = append(
			errorMap["first_name"],
			"First name must be at least 2 characters",
		)
	}
	if r.LastName == "" {
		errorMap["last_name"] = append(errorMap["last_name"], "Last name is required")
	}
	if r.LastName != "" && len(r.LastName) < 2 {
		errorMap["last_name"] = append(
			errorMap["last_name"],
			"Last name must be at least 2 characters",
		)
	}
	return
}
