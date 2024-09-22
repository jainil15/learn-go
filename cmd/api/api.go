package api

import (
	"learn/go/middlewares"
	"learn/go/services/auth"
	"learn/go/services/health"
	"learn/go/services/session"
	"learn/go/services/user"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type IAPIServer interface {
	Run(server *APIServer) error
}

type APIServer struct {
	addr string
	db   *sqlx.DB
}

func NewAPIServer(addr string, db *sqlx.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (server *APIServer) Run() error {
	router := http.NewServeMux()
	port := server.addr
	slog.Info("Server started at", "Port", port)
	// Health handler
	healthHandler := health.NewHandler()
	healthHandler.RegiesterRoutes(router)
	// User handler
	userHandler := user.NewHandler(user.NewStore(server.db))
	userHandler.RegisterRoutes(router)
	// Auth handler
	authHandler := auth.NewHandler(user.NewStore(server.db), session.NewStore(server.db))
	authHandler.RegisterRoutes(router)
	httpServer := http.Server{
		Addr:    server.addr,
		Handler: middlewares.Logging(router),
	}
	return httpServer.ListenAndServe()
}
