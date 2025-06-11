package vaccine

import "time"

type Vaccine struct {
	PatientID      int       `json:"patient_id" validate:"required"`
	VaccineType    string    `json:"vaccine_type" validate:"required"`
	AdministeredOn time.Time `json:"administered_on" validate:"required"`
	Dose           string    `json:"dose" validate:"required"`
	Observations   string    `json:"observations"`
}

type VaccineBase struct {
	VaccineID      int       `json:"vaccine_id"`
	VaccineType    string    `json:"vaccine_type" validate:"required"`
	AdministeredOn time.Time `json:"administered_on" validate:"required"`
	Dose           string    `json:"dose" validate:"required"`
	Observations   string    `json:"observations"`
}
