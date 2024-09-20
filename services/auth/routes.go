package auth

import (
	"learn/go/services/user"
	"learn/go/utils"
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
	user := h.store.GetAll()
	err := utils.ResponseHandler(w, &utils.SuccessResponse{
		StatusCode: http.StatusOK,
		Result:     user,
		Message:    "Success",
	})
	if err != nil {
		return
	}
}
