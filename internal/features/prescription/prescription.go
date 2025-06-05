package prescription

import "time"

type PrescriptionItem struct {
	PrescriptionID      int    `json:"prescription_id"`
	Medication          string `json:"medication"`
	Dosage              string `json:"dosage"`
	Frequency           string `json:"frequency"`
	DurationDays        int    `json:"duration_days"`
	AdministrationRoute string `json:"administration_route"`
	Observations        string `json:"observations"`
}

type Prescription struct {
	PatientID           int                `json:"patient_id"`
	DoctorID            int                `json:"doctor_id"`
	ElectronicSignature string             `json:"electronic_signature"`
	Observations        string             `json:"observations"`
	Items               []PrescriptionItem `json:"items"`
}

type PrescriptionItemBase struct {
	ItemID              int    `json:"item_id"`
	Medication          string `json:"medication"`
	Dosage              string `json:"dosage"`
	DurationDays        int    `json:"duration_days"`
	AdministrationRoute string `json:"administration_route"`
}

type PrescriptionBase struct {
	PrescriptionID int                    `json:"prescription_id"`
	PatientName    string                 `json:"patient_name"`
	PatientDNI     string                 `json:"patient_dni"`
	IssuedAt       time.Time              `json:"issued_at"`
	Items          []PrescriptionItemBase `json:"items"`
}
