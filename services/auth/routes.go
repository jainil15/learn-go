package auth

import (
	"learn/go/services/session"
	"learn/go/services/user"
	"learn/go/utils"
	"log"
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
	valErrors := payload.validate()
	if valErrors != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    "Validation error",
			Error:      valErrors,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	user, err := h.userStore.GetByEmail(payload.Email)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    "User with this email not found",
			StatusCode: http.StatusNotFound,
			Error: map[string][]string{
				"email": {"User with this email not found"},
			},
		})
		return
	}

	if !utils.ComparePassword(payload.Password, user.PasswordHash) {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: "Invalid password",
			Error: map[string][]string{
				"password": {"Invalid password"},
			},
			StatusCode: http.StatusUnauthorized,
		})
		return
	}
	_, err = h.sessionStore.GetSessionByID(user.Id)
	if err != nil {
		_, err := h.sessionStore.CreateSession(user.Id)
		if err != nil {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				Message: err.Error(),
			})
			return
		}
	}
	log.Printf("Request: %v\n", user.Email)
	accessToken, err := user.Createtoken()
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	err = utils.ResponseHandler(
		w,
		&utils.SuccessResponse{
			StatusCode: http.StatusOK,
			Result: map[string]interface{}{
				"user":         user,
				"access_token": accessToken,
			},
			Message: "Success",
		},
	)
	if err != nil {
		return
	}
}
