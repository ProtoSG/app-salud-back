package prescription

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/prescription",
		middleware.Auth(
			middleware.RequireRoles("DOCTOR")(
				http.HandlerFunc(controller.GetAll),
			),
		),
	).Methods("GET")
	r.Handle("/prescription",
		middleware.Auth(
			middleware.RequireRoles("DOCTOR")(
				http.HandlerFunc(controller.Register),
			),
		),
	).Methods("POST")
}
