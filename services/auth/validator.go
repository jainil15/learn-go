package auth

import (
	"regexp"
	"strings"
)

func (payload *LoginPayload) validate() (errorMap map[string][]string) {
	errorMap = make(map[string][]string)
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if strings.TrimSpace(payload.Email) == "" {
		errorMap["email"] = append(errorMap["email"], "Email is required")
	}
	if !emailRegex.MatchString(strings.TrimSpace(payload.Email)) {
		errorMap["email"] = append(errorMap["email"], "Invalid email format")
	}
	if strings.TrimSpace(payload.Password) == "" {
		errorMap["password"] = append(errorMap["password"], "Password is required")
	}
	if len(errorMap) == 0 {
		return nil
	}
	return
}
