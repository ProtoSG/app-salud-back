package medicalappointment

import "time"

type MedicalAppointmentByDoctor struct {
	AppointmentID   int       `json:"appointment_id"`
	PatientID       int       `json:"patient_id"`
	PatientName     string    `json:"patient_name"`
	AppointmentTime time.Time `json:"appintment_time"`
	Status          string    `json:"status"`
	Reason          string    `json:"reason"`
}
