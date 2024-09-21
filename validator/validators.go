package validator

import (
	"regexp"
	"strings"
)

func IsRequired(value string) bool {
	return strings.TrimSpace(value) != ""
}

func IsMinLength(value string, length int) bool {
	return len(strings.TrimSpace(value)) > length
}

func IsMaxLength(value string, length int) bool {
	return len(strings.TrimSpace(value)) < length
}

func IsEmail(value string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(strings.TrimSpace(value)) && strings.TrimSpace(value) != ""
}

func IsRegexMatch(value string, regex string) bool {
	r := regexp.MustCompile(regex)
	return r.MatchString(strings.TrimSpace(value)) && strings.TrimSpace(value) != ""
}
