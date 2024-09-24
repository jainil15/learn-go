package auth

import (
	"learn/go/services/session"
	"learn/go/services/user"
	"net/http"
)

type Handler struct {
	userStore    user.UserStore
	sessionStore session.SessionStore
}

func NewHandler(userStore user.UserStore, sessionStore session.SessionStore) *Handler {
	return &Handler{userStore: userStore, sessionStore: sessionStore}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /auth/login", h.handleLogin)
}
