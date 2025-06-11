package patient

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/patient",
		middleware.Auth(
			middleware.RequireRoles("DOCTOR", "ENFERMERO")(
				http.HandlerFunc(controller.Create),
			),
		),
	).Methods("POST")
	r.Handle("/patient",
		middleware.Auth(
			middleware.RequireRoles("DOCTOR", "ENFERMERO")(
				http.HandlerFunc(controller.GetAll),
			),
		),
	).Methods("GET")
	r.Handle("/patient/{id}",
		middleware.Auth(
			middleware.RequireRoles("DOCTOR", "ENFERMERO")(
				http.HandlerFunc(controller.GetByID),
			),
		),
	).Methods("GET")
}
