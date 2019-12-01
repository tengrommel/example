package api

import (
	v1 "awesomeProject/finance-app-backend/backend/internal/api/v1"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter provide a handler API service
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil
}
