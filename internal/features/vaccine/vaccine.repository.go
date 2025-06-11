package vaccine

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(vaccine *Vaccine) (int, error)
	ReadByPatientID(id int) ([]*VaccineBase, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (p *postgreRepo) Create(vaccine *Vaccine) (int, error) {
	const q = `
		INSERT INTO vaccine (
			patient_id, vaccine_type, administered_on, dose, observations
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING vaccine_id;
	`

	var id int
	if err := p.db.QueryRow(
		q,
		vaccine.PatientID,
		vaccine.VaccineType,
		vaccine.AdministeredOn,
		vaccine.Dose,
		vaccine.Observations,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("error en la consulta: %v", err)
	}

	return id, nil
}

func (p *postgreRepo) ReadByPatientID(id int) ([]*VaccineBase, error) {
	const q = `
		SELECT 
			vaccine_id,
			vaccine_type,
			administered_on,
			dose,
			observations
		FROM vaccine 
		WHERE patient_id = $1
		ORDER BY administered_on ASC;
	`

	vaccines := []*VaccineBase{}
	rows, err := p.db.Query(q, id)
	if err != nil {
		return nil, fmt.Errorf("error en la consulta: %v", err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre las filas: %v", err)
	}

	for rows.Next() {
		vaccine := &VaccineBase{}
		if err := rows.Scan(
			&vaccine.VaccineID,
			&vaccine.VaccineType,
			&vaccine.AdministeredOn,
			&vaccine.Dose,
			&vaccine.Observations,
		); err != nil {
			return nil, fmt.Errorf("error en el escaneo de la fila: %v", err)
		}

		vaccines = append(vaccines, vaccine)
	}

	return vaccines, nil
}
