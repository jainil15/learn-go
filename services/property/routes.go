package property

import (
	"learn/go/middlewares"
	"learn/go/models"
	"learn/go/utils"
	"net/http"
)

type Handler struct {
	propertyStore       models.PropertyStore
	propertyAccessStore models.PropertyAccessStore
}

func NewHandler(
	propertyStore models.PropertyStore,
	propertyAccessStore models.PropertyAccessStore,
) *Handler {
	return &Handler{
		propertyStore:       propertyStore,
		propertyAccessStore: propertyAccessStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /property", middlewares.CheckAccessToken(h.handleCreate))
	router.HandleFunc("GET /property", middlewares.CheckAccessToken(h.handleGetAllByUserId))
}

func (h *Handler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var payload models.CreatePropertyPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid JSON format",
		})
		return
	}

	if err := payload.Validate(); err != nil {

		utils.ErrorHandler(w, &utils.ErrorResponse{
			Error:      err,
			Message:    "Validation Error",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	property, err := h.propertyStore.Create(&payload)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Error:      err,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Error:      err,
			Message:    "Bad user",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	propertyaccess, err := h.propertyAccessStore.Create(property.Id, user.Id)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Error:      err,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	err = utils.ResponseHandler(w, &utils.SuccessResponse{
		StatusCode: http.StatusOK,
		Result:     map[string]interface{}{"property": property, "propertyaccess": propertyaccess},
		Message:    "Success",
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}

func (h *Handler) handleGetAllByUserId(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    "Bad user",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	propertyaccesses, err := h.propertyAccessStore.GetAllByUserId(user.Id)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	err = utils.ResponseHandler(w, &utils.SuccessResponse{
		StatusCode: http.StatusOK,
		Result:     map[string]interface{}{"propertyaccesses": propertyaccesses},
		Message:    "Success",
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}
