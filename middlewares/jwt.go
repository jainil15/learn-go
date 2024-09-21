package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"learn/go/config"
	"learn/go/types"
	"learn/go/utils"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func CheckAccessToken(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessTokenString := getAccessTokenFromRequest(r)
		if accessTokenString == "" {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Missing Access Token",
			})
			return
		}
		accessToken, err := validateAccessToken(accessTokenString)
		if err != nil {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Access Token Malformed or Expired",
			})
			return
		}
		if !accessToken.Valid {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid Token",
			})
			return
		}
		ctx := r.Context()

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid Claims",
			})
			return
		}
		encode, _ := json.Marshal(claims["user"])
		// This is dogshit code maybe...........
		u := types.User{}
		_ = json.Unmarshal(encode, &u)
		ctx = context.WithValue(ctx, "user", u)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}

func getAccessTokenFromRequest(r *http.Request) string {
	accessToken := r.Header.Get("Authorization")
	accessTokenSplit := strings.Split(strings.TrimSpace(accessToken), " ")
	if len(accessTokenSplit) < 2 {
		return ""
	}
	log.Printf("Auth Token :: %v\n\n", accessToken)
	return strings.TrimSpace(strings.Split(accessToken, " ")[1])
}

func validateAccessToken(accessToken string) (*jwt.Token, error) {
	return jwt.Parse(
		accessToken,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected error")
			}
			return []byte(config.Envs.JwtSecret), nil
		},
	)
}
