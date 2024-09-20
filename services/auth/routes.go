package auth

import (
	"learn/go/services/user"
	"net/http"
)

type Handler struct {
	store user.UserStore
}

func NewHandler(store user.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /user/login", h.handleLogin)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}
