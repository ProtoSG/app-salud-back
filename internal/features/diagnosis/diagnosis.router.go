package diagnosis

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/diagnosis",
		middleware.Auth(
			http.HandlerFunc(controller.Register),
		),
	).Methods("POST")
	r.Handle("/diagnosis/{id}",
		middleware.Auth(
			http.HandlerFunc(controller.GetByID),
		),
	).Methods("GET")
}
