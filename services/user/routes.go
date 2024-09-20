package user

import (
	"learn/go/types"
	"learn/go/utils"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	store UserStore
}

func NewHandler(store UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /user", h.handleGetAll)
	router.HandleFunc("POST /user/register", h.handleRegister)
}

func (h *Handler) handleGetAll(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	log.Println(payload)
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(
		payload.Password,
	), 10)
	user, err := h.store.Register(&RegisterUser{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PasswordHash: string(PasswordHash),
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if user == nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{Message: "Error creating user"})
		return
	}
	err = utils.ResponseHandler(w, &utils.SuccessResponse{
		StatusCode: http.StatusOK,
		Result:     user,
		Message:    "Success",
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}
