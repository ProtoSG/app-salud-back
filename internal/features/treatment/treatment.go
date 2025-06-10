package treatment

import "time"

type Treatment struct {
	PatientID    int       `json:"patient_id" validate:"required"`
	DoctorID     int       `json:"doctor_id" validate:"required"`
	StartDate    time.Time `json:"start_date" validate:"required"`
	EndDate      time.Time `json:"end_date" validate:"required"`
	Description  string    `json:"Description" validate:"required"`
	Observations string    `json:"observations" validate:"required"`
}

type TreatmentBase struct {
	TreatmentID  int       `json:"treatment_id"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Description  string    `json:"Description"`
	Observations string    `json:"observations"`
}
