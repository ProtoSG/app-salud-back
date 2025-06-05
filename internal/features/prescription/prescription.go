package prescription

type PrescriptionItem struct {
	ItemID              int    `json:"item_id"`
	PrescriptionID      int    `json:"prescription_id"`
	Medication          string `json:"medication"`
	Dosage              string `json:"dosage"`
	Frequency           string `json:"frequency"`
	DurationDays        int    `json:"duration_days"`
	AdministrationRoute string `json:"administration_route"`
	Observations        string `json:"observations"`
}

type Prescription struct {
	PrescriptionID      int                `json:"prescription_id"`
	PatientID           int                `json:"patient_id"`
	PatientName         string             `json:"patient_name"`
	PatientIdent        string             `json:"patient_identification"`
	DoctorID            int                `json:"doctor_id"`
	ElectronicSignature string             `json:"electronic_signature"`
	Observations        string             `json:"observations"`
	Items               []PrescriptionItem `json:"items"`
}
