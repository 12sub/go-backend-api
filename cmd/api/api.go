package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-rest-api/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	// Initializing routers, creating subrouters and multiplexing the routers
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Creating a new service to register the subrouter
	userStore := user.NewStore(s.db)
	userServices := user.NewHandles(userStore)
	userServices.RegisterRoutes(subrouter)

	// Return Route address
	log.Println("Now listening at: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
