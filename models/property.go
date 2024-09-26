package models

import (
	"learn/go/validator"
	"time"

	"github.com/jmoiron/sqlx"
)

type PropertyStore interface {
	GetById(Id string) (*Property, error)
	Create(*CreatePropertyPayload, *sqlx.Tx) (*Property, error)
}

type Property struct {
	Id          string    `json:"id"           db:"id"`
	Name        string    `json:"name"         db:"name"`
	Email       string    `json:"email"        db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Address     string    `json:"address"      db:"address"`
	About       *string   `json:"about"        db:"about"`
	CreatedAt   time.Time `json:"created_at"   db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"   db:"updated_at"`
}

type CreatePropertyPayload struct {
	Name        string  `json:"name"         db:"name"`
	Email       string  `json:"email"        db:"email"`
	PhoneNumber string  `json:"phone_number" db:"phone_number"`
	Address     string  `json:"address"      db:"address"`
	About       *string `json:"about"        db:"about"`
}

func (c *CreatePropertyPayload) Validate() (errorMap validator.ValidationError) {
	phoneRegex := `^(\+\d{1,2}\s?)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`
	errorMap = validator.ValidationError{}
	if !validator.IsRequired(c.Email) {
		errorMap.Add("email", "Email is required")
	} else if !validator.IsEmail(c.Email) {
		errorMap.Add("email", "Email is not valid")
	}
	if !validator.IsRequired(c.Name) {
		errorMap.Add("name", "Name is required")
	}
	if !validator.IsRequired(c.PhoneNumber) {
		errorMap.Add("phone_number", "Phone number is required")
	} else if !validator.IsRegexMatch(c.PhoneNumber, phoneRegex) {
		errorMap.Add("phone_number", "Phone number is not valid")
	}
	if !validator.IsRequired(c.Address) {
		errorMap.Add("address", "Address is required")
	}
	if c.About != nil && !validator.IsMinLength(*c.About, 1) {
		errorMap.Add("about", "Min 1 length required")
	}
	if errorMap.IsEmpty() {
		return nil
	}
	return
}
