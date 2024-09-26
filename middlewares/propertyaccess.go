package middlewares

import (
	"context"
	"learn/go/models"
	"learn/go/utils"
	"net/http"
)

func CheckPropertyAccess(ps models.PropertyAccessStore, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		propertyId := r.PathValue("propertyId")
		if propertyId == "" {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				Message:    "Property Id required",
				StatusCode: http.StatusForbidden,
			})
			return
		}
		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				Message:    "User context missing",
				StatusCode: http.StatusForbidden,
			})
			return
		}

		property, err := ps.GetByUserIdAndPropertyId(user.Id, propertyId)
		if err != nil {
			utils.ErrorHandler(w, &utils.ErrorResponse{
				Message:    "Property access denied",
				StatusCode: http.StatusForbidden,
			})
			return
		}
		c := context.WithValue(r.Context(), "property", property)
		r = r.WithContext(c)
		h(w, r)
	}
}
