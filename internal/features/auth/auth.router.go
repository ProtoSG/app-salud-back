package auth

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.Handle("/validate", middleware.Auth(
		http.HandlerFunc(controller.Validate),
	)).Methods("GET")
}
