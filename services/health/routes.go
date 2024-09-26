package health

import (
	"learn/go/utils"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegiesterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		err := utils.ResponseHandler(w, &utils.SuccessResponse{
			StatusCode: 200,
			Message:    "Server is running",
		})
		if err != nil {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      struct{}{},
				Message:    "Internal Server Error",
			})
		}
	})
}
