package api

import (
	"learn/go/middlewares"
	"learn/go/routes"
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
	chain := middlewares.CreateStack(middlewares.Logging)
	routes.AddRoutes(router, server.db)
	httpServer := http.Server{
		Addr:    server.addr,
		Handler: chain(router),
	}
	return httpServer.ListenAndServe()
}
