package medicalhistory

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(mh *MedicalHistory) (int, error)
	ReadByPatientID(id int) ([]*MedicalHistoryBase, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (p *postgreRepo) Create(mh *MedicalHistory) (int, error) {
	const q = `
		INSERT INTO medical_history (
			patient_id, description
		) VALUES($1, $2)
		RETURNING history_id;
	`

	var id int
	if err := p.db.QueryRow(
		q,
		mh.PatientID,
		mh.Description,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("error en la consulta: %v", err)
	}

	return id, nil
}

func (p *postgreRepo) ReadByPatientID(id int) ([]*MedicalHistoryBase, error) {
	const q = `
		SELECT 
			history_id,
			description,
		recorded_at
		FROM medical_history
		WHERE history_id = $1
		ORDER BY recorded_at ASC;
		`

	histories := []*MedicalHistoryBase{}
	rows, err := p.db.Query(q, id)
	if err != nil {
		return nil, fmt.Errorf("error en la consulta: %v", err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre las filas: %v", err)
	}

	for rows.Next() {
		history := &MedicalHistoryBase{}
		if err := rows.Scan(
			&history.MedicalHistoryID,
			&history.Description,
			&history.RecordedAt,
		); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}

		histories = append(histories, history)
	}

	return histories, nil
}
