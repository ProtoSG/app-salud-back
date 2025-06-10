package diagnosis

import "time"

type Diagnosis struct {
	PatientID     int       `json:"patient_id" validate:"required"`
	DoctorID      int       `json:"doctor_id" validate:"required"`
	DiagnosisDate time.Time `json:"diagnosis_date" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	Observations  string    `json:"observations"`
}

type DiagnosisBase struct {
	DiagnosisID   int       `json:"diagosis_id"`
	DiagnosisDate time.Time `json:"diagnosis_date"`
	Description   string    `json:"description"`
	Observations  string    `json:"observations"`
}
