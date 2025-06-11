package medicalappointment

import (
	"database/sql"
	"fmt"
	"time"
)

type Repository interface {
	ReadAppointmentsToday(doctor_id int) ([]*MedicalAppointmentByDoctor, error)
	Create(appointment *MedicalAppointment) (int, error)
	ReadByDateRange(start time.Time, end time.Time) ([]*MedicalAppointmentEvent, error)
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

func (p *postgreRepo) Create(appointment *MedicalAppointment) (int, error) {
	const q = `
		INSERT INTO medical_appointment (
			patient_id, doctor_id, appointment_time, duration_minutes, reason
		) VALUES($1, $2, $3, $4, $5)
		RETURNING appointment_id;
	`

	var id int
	if err := p.db.QueryRow(
		q,
		appointment.PatientID,
		appointment.DoctorID,
		appointment.AppointmentTime,
		appointment.DurationMinutes,
		appointment.Reason,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("error en la query: %v", err)
	}

	return id, nil
}

func (p *postgreRepo) ReadByDateRange(start time.Time, end time.Time) ([]*MedicalAppointmentEvent, error) {
	const q = `
		SELECT 
			ma.appointment_id,
			p.first_name || ' ' || p.last_name   AS patient_name,
			d.first_name || ' ' || d.last_name   AS doctor_name,
			ma.appointment_time,
			ma.appointment_time + (ma.duration_minutes * INTERVAL '1 minute')::interval AS end_time,
			ma.reason
		FROM medical_appointment ma
		JOIN patient p ON p.patient_id = ma.patient_id
		JOIN users d ON d.user_id = ma.doctor_id
		WHERE ma.appointment_time 
		BETWEEN $1::timestamptz
		AND $2::timestamptz;
	`

	rows, err := p.db.Query(q, start, end)
	if err != nil {
		return nil, fmt.Errorf("error en la consulta: %v", err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar las filas: %v", err)
	}

	appointments := []*MedicalAppointmentEvent{}
	for rows.Next() {
		appoint := &MedicalAppointmentEvent{}
		if err := rows.Scan(
			&appoint.AppointmentID,
			&appoint.PatientName,
			&appoint.DoctorName,
			&appoint.StartTime,
			&appoint.EndTime,
			&appoint.Reason,
		); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}

		appointments = append(appointments, appoint)
	}

	return appointments, nil
}
