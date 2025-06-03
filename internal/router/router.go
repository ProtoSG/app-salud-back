package router

import (
	"github.com/ProtoSG/app-salud-back/internal/features/auth"
	medicalappointment "github.com/ProtoSG/app-salud-back/internal/features/medicalAppointment"
	"github.com/ProtoSG/app-salud-back/internal/features/usuario"
	"github.com/ProtoSG/app-salud-back/internal/services"
	"github.com/gorilla/mux"
)

func NewRouterContainer(r *mux.Router, svc *services.ServiceContainer) {
	usuario.NewRouter(r, svc.Usuario)
	auth.NewRouter(r, svc.Auth)
	medicalappointment.NewRouter(r, svc.MedicalAppointment)
}
