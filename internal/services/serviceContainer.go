package services

import (
	"database/sql"

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
)

type ServiceContainer struct {
	User               *user.Service
	Auth               *auth.Service
	MedicalAppointment *medicalappointment.Service
	Patient            *patient.Service
	Prescription       *prescription.Service
	Diagnosis          *diagnosis.Service
	Treatment          *treatment.Service
	LabResult          *labresult.Service
	Vaccine            *vaccine.Service
	MedicalHistory     *medicalhistory.Service
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	userRepo := user.NewPostgreRepo(db)
	authRepo := auth.NewPostgreRepo(db)
	mAppt := medicalappointment.NewPostgreRepo(db)
	patientRepo := patient.NewPostgreRepo(db)
	presRepo := prescription.NewPostgreRepo(db)
	diagnosisRepo := diagnosis.NewPostgreRepo(db)
	treatmentRepo := treatment.NewPostgreRepo(db)
	labResultRepo := labresult.NewPostgreRepository(db)
	vaccineRepo := vaccine.NewPostgreRepo(db)
	mhRepo := medicalhistory.NewPostgreRepo(db)

	return &ServiceContainer{
		User:               user.NewService(userRepo),
		Auth:               auth.NewService(authRepo),
		MedicalAppointment: medicalappointment.NewService(mAppt),
		Patient:            patient.NewService(patientRepo),
		Prescription:       prescription.NewService(presRepo),
		Diagnosis:          diagnosis.NewService(diagnosisRepo),
		Treatment:          treatment.NewService(treatmentRepo),
		LabResult:          labresult.NewService(labResultRepo),
		Vaccine:            vaccine.NewService(vaccineRepo),
		MedicalHistory:     medicalhistory.NewService(mhRepo),
	}
}
