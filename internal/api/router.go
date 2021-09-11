package api

import (
	"net/http"

	v1 "github.com/Sunshine31/financial-app/internal/api/v1"
	"github.com/gorilla/mux"
)

func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil
}
