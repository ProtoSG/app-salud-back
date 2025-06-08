package patient

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Repository interface {
	Create(patient *Patient) (int64, error)
	FindPatientByDNI(dni string) error
	Read(page, limit int, filters PatientFilters) ([]*PatientBasicData, error)
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

func (this *postgreRepo) Read(
	page, limit int,
	filters PatientFilters,
) ([]*PatientBasicData, error) {
	offset := (page - 1) * limit

	whereParts := []string{"true"}
	args := []any{}
	argPos := 1

	if filters.Gender != "" {
		whereParts = append(whereParts,
			fmt.Sprintf("gender = $%d", argPos),
		)
		args = append(args, filters.Gender)
		argPos++
	}

	if filters.RangeAge[0] != 0 || filters.RangeAge[1] != 0 {
		if filters.RangeAge[0] != 0 && filters.RangeAge[1] != 0 {
			whereParts = append(whereParts,
				fmt.Sprintf("age BETWEEN $%d AND $%d", argPos, argPos+1),
			)
			args = append(args,
				filters.RangeAge[0],
				filters.RangeAge[1],
			)
			argPos += 2
		} else if filters.RangeAge[0] != 0 {
			whereParts = append(whereParts,
				fmt.Sprintf("age >= $%d", argPos),
			)
			args = append(args, filters.RangeAge[0])
			argPos++
		} else {
			whereParts = append(whereParts,
				fmt.Sprintf("age <= $%d", argPos),
			)
			args = append(args, filters.RangeAge[1])
			argPos++
		}
	}

	if filters.Disease != "" {
		whereParts = append(whereParts,
			fmt.Sprintf("diseases @> ARRAY[$%d]::varchar[]", argPos),
		)
		args = append(args, filters.Disease)
		argPos++
	}

	whereSQL := strings.Join(whereParts, " AND ")
	args = append(args, limit, offset)
	limPos := argPos
	offPos := argPos + 1

	q := fmt.Sprintf(`
		WITH patient_info AS (
			SELECT
				p.patient_id,
				p.first_name || ' ' || p.last_name AS full_name,
				p.gender,
				EXTRACT(YEAR FROM AGE(CURRENT_DATE, p.birth_date))::INT AS age,
				array_agg(d.description) AS diseases
			FROM patient p
			JOIN diagnosis d
				ON d.patient_id = p.patient_id
			WHERE NOT p.is_deleted
			GROUP BY
				p.patient_id, p.first_name, p.last_name, p.gender, p.birth_date
		)
		SELECT
			patient_id,
			full_name,
			gender,
			age
		FROM patient_info
		WHERE %s
		LIMIT $%d
		OFFSET $%d;
		`, whereSQL, limPos, offPos)

	rows, err := this.db.Query(q, args...)
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
