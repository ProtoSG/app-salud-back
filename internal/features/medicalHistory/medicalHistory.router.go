package medicalhistory

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/medicalhistory", middleware.Auth(
		http.HandlerFunc(controller.Register),
	)).Methods("POST")

	r.Handle("/medicalhistory/patient/{id}", middleware.Auth(
		http.HandlerFunc(controller.GetByPatientID),
	)).Methods("GET")
}
