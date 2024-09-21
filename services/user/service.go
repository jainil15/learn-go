package user

import (
	"fmt"
	"learn/go/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (u *User) Createtoken() (string, error) {
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
