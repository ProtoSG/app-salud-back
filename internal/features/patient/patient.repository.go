package patient

import (
	"database/sql"
	"fmt"

	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Repository interface {
	Create(patient *Patient) (int64, error)
	FindPatientByDNI(dni string) error
	Read() ([]*PatientBasicData, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (this *postgreRepo) Create(patient *Patient) (int64, error) {
	const q = `
		INSERT INTO patient 
		(first_name,
			last_name,
			dni,
			birth_date,
			gender,
			address,
			phone,
			email,
			photo_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING patient_id;
	`

	var patient_id int64
	if err := this.db.QueryRow(q,
		patient.FirstName,
		patient.LastName,
		patient.DNI,
		patient.BirthDate,
		patient.Gender,
		patient.Address,
		patient.Phone,
		patient.Email,
		patient.PhotoURL,
	).Scan(&patient_id); err != nil {
		return 0, fmt.Errorf("Error al crear el paciente")
	}

	return patient_id, nil
}

func (this *postgreRepo) FindPatientByDNI(dni string) error {
	const q = `
		SELECT EXISTS (
			SELECT 1
			FROM patient
			WHERE is_deleted = FALSE
			AND dni = $1
		);
	`
	var exists bool
	err := this.db.QueryRow(q, dni).Scan(&exists)
	if err != nil {
		return fmt.Errorf("Error en la consulta: %w", err)
	}

	if !exists {
		return utils.NewEntityNotFound(dni, "DNI", "Patient")
	}

	return nil
}

func (this *postgreRepo) Read() ([]*PatientBasicData, error) {
	const q = `
		SELECT
			p.patient_id,
			(p.first_name || ' ' || p.last_name) AS full_name,
			p.gender,
			EXTRACT(YEAR FROM AGE(CURRENT_DATE, p.birth_date))::INT
		FROM patient p
		WHERE p.is_deleted = FALSE 
	`

	rows, err := this.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("Error en la consulta: %w", err)
	}
	defer rows.Close()

	patients := []*PatientBasicData{}
	for rows.Next() {
		patient := &PatientBasicData{}
		if err := rows.Scan(
			&patient.PatientID,
			&patient.FullName,
			&patient.Gender,
			&patient.Age,
		); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %w", err)
		}
		patients = append(patients, patient)
	}

	return patients, nil
}
