package api

import (
	"learn/go/services/auth"
	"learn/go/services/health"
	"learn/go/services/user"
	"log"
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
	log.Println("Server started at 8080 ")
	// Health handler
	healthHandler := health.NewHandler()
	healthHandler.RegiesterRoutes(router)
	// User handler
	userHandler := user.NewHandler(user.NewStore(server.db))
	userHandler.RegisterRoutes(router)
	// Auth handler
	authHandler := auth.NewHandler(user.NewStore(server.db))
	authHandler.RegisterRoutes(router)

	return http.ListenAndServe(server.addr, router)
}
