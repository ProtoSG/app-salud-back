package treatment

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(treatment *Treatment) (int, error)
	ReadByPatientID(id int) ([]*TreatmentBase, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (r *postgreRepo) Create(treatment *Treatment) (int, error) {
	const q = `
		INSERT INTO treatment (
			patient_id, doctor_id, start_date, end_date, description, observations
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING treatment_id;
	`

	var id int
	if err := r.db.QueryRow(q,
		treatment.PatientID,
		treatment.DoctorID,
		treatment.StartDate,
		treatment.EndDate,
		treatment.Description,
		treatment.Observations,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("error en la consulta: %v", err)
	}

	return id, nil
}

func (r *postgreRepo) ReadByPatientID(id int) ([]*TreatmentBase, error) {
	const q = `
		SELECT treatment_id, start_date, end_date, description, observations
		FROM treatment
		WHERE patient_id = $1
		ORDER BY start_date ASC;
	`

	treatments := []*TreatmentBase{}
	rows, err := r.db.Query(q, id)
	if err != nil {
		return nil, fmt.Errorf("error en la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var treatment TreatmentBase
		if err := rows.Scan(
			&treatment.TreatmentID,
			&treatment.StartDate,
			&treatment.EndDate,
			&treatment.Description,
			&treatment.Observations,
		); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		treatments = append(treatments, &treatment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre las filas: %v", err)
	}

	return treatments, nil
}
