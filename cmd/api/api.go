package api

import (
	"learn/go/middlewares"
	"learn/go/services/auth"
	"learn/go/services/health"
	"learn/go/services/session"
	"learn/go/services/user"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type IAPIServer interface {
	Run(server *APIServer) error
}

type APIServer struct {
	addr   string
	db     *sqlx.DB
	logger *middlewares.WrappedLogger
}

func NewAPIServer(addr string, db *sqlx.DB) *APIServer {
	return &APIServer{
		addr:   addr,
		db:     db,
		logger: &middlewares.WrappedLogger{},
	}
}

func (server *APIServer) Run() error {
	router := http.NewServeMux()
	port := server.addr
	server.logger.Debug("Server started at ", "Port", port)
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
