package router

import (
	"github.com/ProtoSG/app-salud-back/internal/features/auth"
	medicalappointment "github.com/ProtoSG/app-salud-back/internal/features/medicalAppointment"
	"github.com/ProtoSG/app-salud-back/internal/features/patient"
	"github.com/ProtoSG/app-salud-back/internal/features/user"
	"github.com/ProtoSG/app-salud-back/internal/services"
	"github.com/gorilla/mux"
)

func NewRouterContainer(r *mux.Router, svc *services.ServiceContainer) {
	user.NewRouter(r, svc.User)
	auth.NewRouter(r, svc.Auth)
	medicalappointment.NewRouter(r, svc.MedicalAppointment)
	patient.NewRouter(r, svc.Patient)
}
