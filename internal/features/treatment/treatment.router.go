package treatment

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/treatment",
		middleware.Auth(
			http.HandlerFunc(controller.CreateTreatment),
		),
	).Methods("POST")
	r.Handle("/treatment/patient{id}",
		middleware.Auth(
			http.HandlerFunc(controller.GetTreatmentsByPatientID),
		),
	).Methods("GET")
}
