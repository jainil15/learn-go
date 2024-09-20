package user

import (
	"learn/go/config"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func (r *RegisterUserPayload) validate() (errorMap map[string][]string) {
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
	if len(r.FirstName) < 2 {
		errorMap["first_name"] = append(
			errorMap["first_name"],
			"First name must be at least 2 characters",
		)
	}
	if r.LastName == "" {
		errorMap["last_name"] = append(errorMap["last_name"], "Last name is required")
	}
	if len(r.LastName) < 2 {
		errorMap["last_name"] = append(
			errorMap["last_name"],
			"Last name must be at least 2 characters",
		)
	}
	return
}

func (l *LoginUser) validate() (errorMap map[string]interface{}) {
	err := config.Validate.Struct(l)
	if err != nil {
		return map[string]interface{}{
			"error": err.(validator.ValidationErrors),
		}
	}
	return nil
}
