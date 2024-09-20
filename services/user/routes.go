package user

import (
	"learn/go/types"
	"learn/go/utils"
	"net/http"
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
	registerUserPayload := &RegisterUserPayload{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	}
	error := registerUserPayload.validate()
	if len(error) > 0 {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    "Validation error",
			Error:      error,
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	if existingUser, _ := h.store.GetByEmail(payload.Email); existingUser != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    "User with this email already exists",
			StatusCode: http.StatusConflict,
			Error: map[string]interface{}{
				"email": []string{"User with this email already exists"},
			},
		})
		return
	}

	PasswordHash, err := utils.EncryptPassword(payload.Password)
	user, err := h.store.Register(&RegisterUser{
		FirstName:    registerUserPayload.FirstName,
		LastName:     registerUserPayload.LastName,
		Email:        registerUserPayload.Email,
		PasswordHash: string(PasswordHash),
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: "Error creating user",
			Error:   err.Error(),
		})
		return
	}
	if user == nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Error:   map[string]interface{}{"server": "Error creating user"},
			Message: "Error creating user",
		})
		return
	}
	err = utils.ResponseHandler(w, &utils.SuccessResponse{
		StatusCode: http.StatusOK,
		Result:     map[string]interface{}{"user": user},
		Message:    "Success",
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}
