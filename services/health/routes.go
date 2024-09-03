package health

import (
	"learn/go/utils"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegiesterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /server-health", func(w http.ResponseWriter, r *http.Request) {

		err := utils.ResponseHandler(w, &utils.SuccessResponse{
			StatusCode: 200,
			Result: struct {
				Res string `json:"res"`
			}{
				Res: "Hello World",
			},
			Message: "Success",
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
