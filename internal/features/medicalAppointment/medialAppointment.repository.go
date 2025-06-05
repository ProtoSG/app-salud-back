package medicalappointment

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	ReadAppointmentsToday(doctor_id int) ([]*MedicalAppointmentByDoctor, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (this *postgreRepo) ReadAppointmentsToday(doctor_id int) ([]*MedicalAppointmentByDoctor, error) {
	const q = `
		SELECT 
			ma.appointment_id,
			ma.patient_id,
			(p.first_name || ' ' || p.last_name) AS patient_name,
			ma.appointment_time,
			ma.status,
			ma.reason
		FROM medical_appointment ma 
		INNER JOIN patient p ON p.patient_id = ma.patient_id
		WHERE ma.doctor_id = $1
  	AND (ma.appointment_time AT TIME ZONE 'America/Lima')::date = CURRENT_DATE
		ORDER BY appointment_time ASC;
	`
	appointments := []*MedicalAppointmentByDoctor{}

	rows, err := this.db.Query(q, doctor_id)
	if err != nil {
		return nil, fmt.Errorf("Error en la consulta: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		appointment := &MedicalAppointmentByDoctor{}
		if err := rows.Scan(
			&appointment.AppointmentID,
			&appointment.PatientID,
			&appointment.PatientName,
			&appointment.AppointmentTime,
			&appointment.Status,
			&appointment.Reason,
		); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %w", err)
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
