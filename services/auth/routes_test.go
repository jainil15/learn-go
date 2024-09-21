package auth

import (
	"bytes"
	"encoding/json"
	"learn/go/services/user"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthServiceHandler(t *testing.T) {
	mockStore := &user.MockUserStore{}
	handler := NewHandler(mockStore)
	if handler == nil {
		t.Error("Handler is nil")
	}
	t.Run("TEST", func(t *testing.T) {
		payload := &LoginPayload{
			Email:    "jainil115@gmail.com",
			Password: "Passw2ord",
		}
		marshal, err := json.Marshal(payload)
		log.Printf("Marshal: %v", string(marshal))
		if err != nil {
			t.Error("Error in marshalling")
		}
		req, err := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(marshal))
		if err != nil {
			t.Error("Error in creating request")
		}
		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /auth/login", handler.handleLogin)
		router.ServeHTTP(rr, req)
		log.Printf("Response: %v", rr.Code)
		if rr.Code != http.StatusOK {
			t.Errorf("Error in response %v", rr.Body)
		}
	})
}