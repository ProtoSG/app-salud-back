package patient

import (
	"net/http"

	"github.com/ProtoSG/app-salud-back/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/patient", middleware.Auth(http.HandlerFunc(controller.Create))).Methods("POST")
	r.Handle("/patient", middleware.Auth(http.HandlerFunc(controller.GetAll))).Methods("GET")
}
