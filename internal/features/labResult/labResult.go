package labresult

import "time"

type LabResult struct {
	PatientID    int       `json:"patient_id" validate:"required"`
	DoctorID     int       `json:"doctor_id" validate:"required"`
	SampleDate   time.Time `json:"sample_date" validate:"required"`
	TestType     string    `json:"test_type" validate:"required"`
	ResultValue  string    `json:"result_value" validate:"required"`
	Observations string    `json:"observations"`
}

type LabResultBase struct {
	LabResultID  int       `json:"lab_result_id"`
	SampleDate   time.Time `json:"sample_date"`
	TestType     string    `json:"test_type"`
	ResultValue  string    `json:"result_value"`
	Observations string    `json:"observations"`
}
