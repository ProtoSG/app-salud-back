package medicalappointment

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/medicalappointment/today",
		middleware.Auth(
			middleware.RequireRoles("DOCTOR", "ADMINISTRADOR")(
				http.HandlerFunc(controller.GetAppointmentsToday),
			),
		),
	).Methods("GET")

	r.Handle("/medicalappointment", middleware.Auth(
		middleware.RequireRoles("ADMINISTRADOR")(
			http.HandlerFunc(controller.Create),
		),
	)).Methods("POST")

	r.Handle("/medicalappointment", middleware.Auth(
		middleware.RequireRoles("ADMINISTRADOR")(
			http.HandlerFunc(controller.ListRange),
		),
	)).Methods("GET")
}
