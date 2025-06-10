package labresult

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(labResult *LabResult) (int, error)
	ReadByPatientID(id int) ([]*LabResultBase, error)
}

type postgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(db *sql.DB) Repository {
	return &postgreRepository{db}
}

func (p *postgreRepository) Create(labResult *LabResult) (int, error) {
	const q = `
		INSERT INTO lab_result (
			patient_id, doctor_id, sample_date, test_type, result_value, observations
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING lab_result_id;
	`

	var id int
	if err := p.db.QueryRow(
		q,
		labResult.PatientID,
		labResult.DoctorID,
		labResult.SampleDate,
		labResult.TestType,
		labResult.ResultValue,
		labResult.Observations,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("error en la consulta: %v", err)
	}

	return id, nil
}

func (p *postgreRepository) ReadByPatientID(id int) ([]*LabResultBase, error) {
	const q = `
		SELECT 
			lab_result_id,
			sample_date,
			test_type,
			result_value,
			observations
		FROM lab_result
		WHERE patient_id = $1
		ORDER BY sample_date ASC;
	`

	results := []*LabResultBase{}
	rows, err := p.db.Query(q, id)
	if err != nil {
		return nil, fmt.Errorf("error en la consulta: %v", err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error al iterar sobre las filas: %v", err)
	}

	for rows.Next() {
		result := &LabResultBase{}
		if err := rows.Scan(
			&result.LabResultID,
			&result.SampleDate,
			&result.TestType,
			&result.ResultValue,
			&result.Observations,
		); err != nil {
			return nil, fmt.Errorf("error en el escaneo de la fila: %v", err)
		}

		results = append(results, result)
	}

	return results, nil
}
