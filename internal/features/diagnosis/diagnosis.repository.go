package diagnosis

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(diagnosis *Diagnosis) (int, error)
	ReadByID(id int) ([]*DiagnosisBase, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (this *postgreRepo) Create(diagnosis *Diagnosis) (int, error) {
	const q = `
		INSERT INTO diagnosis (
			patient_id, doctor_id, diagnosis_date, description, observations
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING diagnosis_id;
	`

	var id int
	if err := this.db.QueryRow(
		q,
		diagnosis.PatientID,
		diagnosis.DoctorID,
		diagnosis.DiagnosisDate,
		diagnosis.Description,
		diagnosis.Observations,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("Error en la consulta: %v", err)
	}

	return id, nil
}

func (this *postgreRepo) ReadByID(id int) ([]*DiagnosisBase, error) {
	const q = `
		SELECT 
			diagnosis_id,
			diagnosis_date,
			description,
			observations
		FROM diagnosis
		WHERE patient_id = $1
		ORDER BY diagnosis_date ASC;
	`

	results := []*DiagnosisBase{}
	rows, err := this.db.Query(q, id)
	if err != nil {
		return nil, fmt.Errorf("Error en la consulta: %v", err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error al iterar sobre las filas: %v", err)
	}

	for rows.Next() {
		diagnosis := &DiagnosisBase{}
		if err := rows.Scan(
			&diagnosis.DiagnosisID,
			&diagnosis.DiagnosisDate,
			&diagnosis.Description,
			&diagnosis.Observations,
		); err != nil {
			return nil, fmt.Errorf("Error en el escaneo de la fila: %v", err)
		}

		results = append(results, diagnosis)
	}

	return results, nil
}
