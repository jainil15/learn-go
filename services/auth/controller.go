package auth

import (
	"learn/go/models"
	"learn/go/utils"
	"net/http"
)

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	payload := models.LoginPayload{}
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	valErrors := payload.Validate()
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

	accessToken, err := utils.CreateToken(user)
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
