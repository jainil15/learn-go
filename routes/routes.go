package routes

import (
	"learn/go/services/auth"
	"learn/go/services/health"
	"learn/go/services/property"
	"learn/go/services/propertyaccess"
	"learn/go/services/session"
	"learn/go/services/user"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func AddRoutes(router *http.ServeMux, db *sqlx.DB) {
	// Health handler
	healthHandler := health.NewHandler()
	healthHandler.RegiesterRoutes(router)
	// User handler
	userHandler := user.NewHandler(user.NewStore(db))
	userHandler.RegisterRoutes(router)
	// Auth handler
	authHandler := auth.NewHandler(user.NewStore(db), session.NewStore(db))
	authHandler.RegisterRoutes(router)
	// Property handler
	propertyHandler := property.NewHandler(
		property.NewStore(db),
		propertyaccess.NewStore(db),
		db,
	)
	propertyHandler.RegisterRoutes(router)
}
