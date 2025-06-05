package prescription

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.Handle("/prescription", http.HandlerFunc(controller.GetAll)).Methods("GET")
	r.Handle("/prescription", http.HandlerFunc(controller.Register)).Methods("POST")
}
