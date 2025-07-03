package router

import (
	"github.com/ProtoSG/app-salud-back/internal/features/auth"
	"github.com/ProtoSG/app-salud-back/internal/features/diagnosis"
	labresult "github.com/ProtoSG/app-salud-back/internal/features/labResult"
	medicalappointment "github.com/ProtoSG/app-salud-back/internal/features/medicalAppointment"
	medicalhistory "github.com/ProtoSG/app-salud-back/internal/features/medicalHistory"
	"github.com/ProtoSG/app-salud-back/internal/features/patient"
	"github.com/ProtoSG/app-salud-back/internal/features/prescription"
	"github.com/ProtoSG/app-salud-back/internal/features/treatment"
	"github.com/ProtoSG/app-salud-back/internal/features/user"
	"github.com/ProtoSG/app-salud-back/internal/features/vaccine"
	"github.com/ProtoSG/app-salud-back/internal/services"
	"github.com/gorilla/mux"
)

func NewRouterContainer(r *mux.Router, svc *services.ServiceContainer) {
	user.NewRouter(r, svc.User)
	auth.NewRouter(r, svc.Auth)
	medicalappointment.NewRouter(r, svc.MedicalAppointment)
	patient.NewRouter(r, svc.Patient)
	prescription.NewRouter(r, svc.Prescription)
	diagnosis.NewRouter(r, svc.Diagnosis)
	treatment.NewRouter(r, svc.Treatment)
	labresult.NewRouter(r, svc.LabResult)
	medicalhistory.NewRouter(r, svc.MedicalHistory)
	vaccine.NewRouter(r, svc.Vaccine)
}
