package auth

import "github.com/gorilla/mux"

func NewRouter(r *mux.Router, svc *Service) {
	controller := NewController(svc)

	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/register", controller.Register).Methods("POST")
}
