package api

import (
	v1 "awesomeProject/finance-app-backend/backend/internal/api/v1"
	"awesomeProject/finance-app-backend/backend/internal/database"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter provide a handler API service
func NewRouter(db database.Database) (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil
}
