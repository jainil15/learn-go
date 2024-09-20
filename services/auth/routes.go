package auth

import (
	"learn/go/services/user"
	"learn/go/utils"
	"log"
	"net/http"
)

type Handler struct {
	store user.UserStore
}

func NewHandler(store user.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /auth/login", h.handleLogin)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	payload := LoginPayload{}
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	user, err := h.store.GetByEmail(payload.Email)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusNotFound,
		})
		return
	}

	if !utils.ComparePassword(payload.Password, user.PasswordHash) {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    "Invalid password",
			StatusCode: http.StatusUnauthorized,
		})
		return
	}
	log.Printf("Request: %v", user.Email)
	err = utils.ResponseHandler(
		w,
		&utils.SuccessResponse{
			StatusCode: http.StatusOK,
			Result:     user,
			Message:    "Success",
		},
	)
	if err != nil {
		return
	}
}
