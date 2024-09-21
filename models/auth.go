package models

import (
	"learn/go/validator"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (payload *LoginPayload) Validate() (errorMap validator.ValidationError) {
	errorMap = validator.ValidationError{}
	if !validator.IsRequired(payload.Email) {
		errorMap.Add("email", "Email is required")
	}
	if !validator.IsEmail(payload.Email) && validator.IsRequired(payload.Email) {
		errorMap.Add("email", "Email is not valid")
	}
	if !validator.IsRequired(payload.Password) {
		errorMap.Add("password", "Password is required")
	}
	if errorMap.IsEmpty() {
		return nil
	}
	return
}
