package services

import (
	"database/sql"

	"github.com/ProtoSG/app-salud-back/internal/features/auth"
	medicalappointment "github.com/ProtoSG/app-salud-back/internal/features/medicalAppointment"
	"github.com/ProtoSG/app-salud-back/internal/features/patient"
	"github.com/ProtoSG/app-salud-back/internal/features/user"
)

type ServiceContainer struct {
	User               *user.Service
	Auth               *auth.Service
	MedicalAppointment *medicalappointment.Service
	Patient            *patient.Service
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	userRepo := user.NewPostgreRepo(db)
	authRepo := auth.NewPostgreRepo(db)
	mAppt := medicalappointment.NewPostgreRepo(db)
	patientRepo := patient.NewPostgreRepo(db)

	return &ServiceContainer{
		User:               user.NewService(userRepo),
		Auth:               auth.NewService(authRepo),
		MedicalAppointment: medicalappointment.NewService(mAppt),
		Patient:            patient.NewService(patientRepo),
	}
}
