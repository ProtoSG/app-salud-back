package medicalappointment

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/medicalappointment",
		middleware.Auth(
			middleware.RequireRoles("DOCTOR", "ENFERMERO")(
				http.HandlerFunc(controller.GetAppointmentsToday),
			),
		),
	).Methods("GET")
}
