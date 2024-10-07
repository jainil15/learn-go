package models

import (
	"learn/go/validator"
	"log"
	"regexp"
	"strings"
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

func (r *RegisterUserPayload) Validate() (errorMap validator.ValidationError) {
	errorMap = make(validator.ValidationError)
	log.Println("Validation Enter")
	if strings.TrimSpace(r.Email) == "" {
		errorMap.Add("email", "Email is required")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(strings.TrimSpace(r.Email)) && strings.TrimSpace(r.Email) != "" {
		errorMap.Add("email", "Email is not valid")
	}
	if strings.TrimSpace(r.Password) == "" {
		errorMap.Add("password", "Password is required")
	}
	if len(r.Password) < 6 && strings.TrimSpace(r.Password) != "" {
		errorMap.Add("password", "Password must be at least 6 characters")
	}
	if strings.TrimSpace(r.FirstName) == "" {
		errorMap.Add("first_name", "First name is required")
	}
	if strings.TrimSpace(r.FirstName) != "" && len(strings.TrimSpace(r.FirstName)) < 2 {
		errorMap.Add("first_name", "First name must be at least 2 characters")
	}
	if strings.TrimSpace(r.LastName) == "" {
		errorMap.Add("last_name", "Last name is required")
	}
	if strings.TrimSpace(r.LastName) != "" && len(strings.TrimSpace(r.LastName)) < 2 {
		errorMap.Add("last_name", "Last name must be at least 2 characters")
	}
	log.Printf("Error: %v\n", errorMap)
	if len(errorMap) <= 0 {
		return nil
	}
	return errorMap
}
