package utils

import (
	"fmt"
	"learn/go/config"
	"learn/go/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CreateToken(u *models.User) (string, error) {
	claim := jwt.MapClaims{
		"iss":  "learngo",
		"aud":  []string{"user"},
		"sub":  fmt.Sprintf("%s_%s", u.FirstName, u.LastName),
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
		"user": u,
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := claims.SignedString([]byte(config.Envs.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
