package medicalhistory

import "time"

type MedicalHistory struct {
	PatientID   int    `json:"patient_id" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type MedicalHistoryBase struct {
	MedicalHistoryID int       `json:"medical_history_id"`
	Description      string    `json:"description"`
	RecordedAt       time.Time `json:"recorded_at"`
}
