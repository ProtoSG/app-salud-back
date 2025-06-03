package services

import (
	"database/sql"

	"github.com/ProtoSG/app-salud-back/internal/features/auth"
	medicalappointment "github.com/ProtoSG/app-salud-back/internal/features/medicalAppointment"
	"github.com/ProtoSG/app-salud-back/internal/features/usuario"
)

type ServiceContainer struct {
	Usuario            *usuario.Service
	Auth               *auth.Service
	MedicalAppointment *medicalappointment.Service
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	userRepo := usuario.NewPostgreRepo(db)
	authRepo := auth.NewPostgreRepo(db)
	mAppt := medicalappointment.NewPostgreRepo(db)

	return &ServiceContainer{
		Usuario:            usuario.NewService(userRepo),
		Auth:               auth.NewService(authRepo),
		MedicalAppointment: medicalappointment.NewService(mAppt),
	}
}
