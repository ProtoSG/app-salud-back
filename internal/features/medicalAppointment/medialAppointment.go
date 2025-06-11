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

type MedicalAppointment struct {
	PatientID       int       `json:"patient_id" validate:"required"`
	DoctorID        int       `json:"doctor_id" validate:"required"`
	AppointmentTime time.Time `json:"appointment_time" validate:"required"`
	DurationMinutes int       `json:"duration_minutes" validate:"required"`
	Reason          string    `json:"reason" validate:"required"`
}

type MedicalAppointmentEvent struct {
	AppointmentID int       `json:"appointment_id"`
	PatientName   string    `json:"patient_name"`
	DoctorName    string    `json:"doctor_name"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	Reason        string    `json:"reason"`
}
