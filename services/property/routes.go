package property

import (
	"learn/go/middlewares"
	"learn/go/models"
	"learn/go/utils"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	db                  *sqlx.DB
	propertyStore       models.PropertyStore
	propertyAccessStore models.PropertyAccessStore
}

func NewHandler(
	propertyStore models.PropertyStore,
	propertyAccessStore models.PropertyAccessStore,
	db *sqlx.DB,
) *Handler {
	return &Handler{
		propertyStore:       propertyStore,
		propertyAccessStore: propertyAccessStore,
		db:                  db,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /property", middlewares.CheckAccessToken(h.handleGetAllByUserId))
	router.HandleFunc(
		"GET /property/{propertyId}",
		middlewares.CheckAccessToken(
			middlewares.CheckPropertyAccess(h.propertyAccessStore, h.handleGetById),
		),
	)
	router.HandleFunc("POST /property", middlewares.CheckAccessToken(h.handleCreate))
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
	tx, err := h.db.Beginx()
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Error:      err,
			Message:    "Error creating transaction",
			StatusCode: http.StatusInternalServerError,
		})
	}
	defer tx.Rollback()
	property, err := h.propertyStore.Create(&payload, tx)
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
	propertyaccess, err := h.propertyAccessStore.Create(property.Id, user.Id, tx)
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
	tx.Commit()
}

// Handle GET all by User Id
func (h *Handler) handleGetAllByUserId(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    "Bad user",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	properties, err := h.propertyAccessStore.GetAllByUserId(user.Id)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	err = utils.ResponseHandler(w, &utils.SuccessResponse{
		StatusCode: http.StatusOK,
		Result:     map[string]interface{}{"properties": properties},
		Message:    "Success",
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}

// Handle get project by Id
func (h *Handler) handleGetById(w http.ResponseWriter, r *http.Request) {
	propertyId := r.PathValue("propertyId")
	if propertyId == "" {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: "Property Id missing",
			Error: map[string]interface{}{
				"propertyId": []string{"Property id missing"},
			},
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	property, err := h.propertyStore.GetById(propertyId)
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: "Property not found",
			Error: map[string]interface{}{
				"propertyId": []string{"Property not found with property id"},
			},
			StatusCode: http.StatusNotFound,
		})
		return
	}
	err = utils.ResponseHandler(w, &utils.SuccessResponse{
		StatusCode: http.StatusOK,
		Result:     map[string]interface{}{"property": property},
		Message:    "Success",
	})
	if err != nil {
		utils.ErrorHandler(w, &utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}
