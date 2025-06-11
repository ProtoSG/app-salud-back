package medicalappointment

import (
	"fmt"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (this *Service) ReadAppointmentsToday(doctor_id int) ([]*MedicalAppointmentByDoctor, error) {
	return this.repo.ReadAppointmentsToday(doctor_id)
}

func (s *Service) Create(
	patientID,
	doctorID int,
	appointmentTime time.Time,
	durationMinutes int,
	reason string,
) (int, error) {
	appointment := &MedicalAppointment{
		PatientID:       patientID,
		DoctorID:        doctorID,
		AppointmentTime: appointmentTime,
		DurationMinutes: durationMinutes,
		Reason:          reason,
	}

	return s.repo.Create(appointment)
}

func (s *Service) ReadByDateRange(start time.Time, end time.Time) ([]*MedicalAppointmentEvent, error) {
	appointments, err := s.repo.ReadByDateRange(start, end)
	if err != nil {
		return nil, err
	}

	if len(appointments) == 0 {
		return nil, fmt.Errorf("No hay eventos en la fecha")
	}

	return appointments, nil
}
