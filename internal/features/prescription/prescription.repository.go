package prescription

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Repository interface {
	Create(pres *Prescription) (int64, error)
	Read() ([]*PrescriptionBase, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (this *postgreRepo) Create(pres *Prescription) (int64, error) {
	ctx := context.Background()
	tx, err := this.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return 0, fmt.Errorf("Error en la transaccion: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	insertPresQuery := `
		INSERT INTO prescription (
		  patient_id, doctor_id, electronic_signature, observations
		) VALUES ($1, $2, $3, $4)
		RETURNING prescription_id;
	`
	var prescriptionID int64
	err = tx.QueryRowContext(ctx, insertPresQuery,
		pres.PatientID,
		pres.DoctorID,
		pres.ElectronicSignature,
		pres.Observations,
	).Scan(&prescriptionID)
	if err != nil {
		_ = tx.Rollback()
		return 0, fmt.Errorf("error al insertar prescription: %w", err)
	}

	insertItemQuery := `
		INSERT INTO prescription_item (
		  prescription_id,
		  medication,
		  dosage,
		  frequency,
		  duration_days,
		  administration_route,
		  observations
		) VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	for _, item := range pres.Items {
		if _, err := tx.ExecContext(ctx, insertItemQuery,
			prescriptionID,
			item.Medication,
			item.Dosage,
			item.Frequency,
			item.DurationDays,
			item.AdministrationRoute,
			item.Observations,
		); err != nil {
			_ = tx.Rollback()
			return 0, fmt.Errorf("Error al insertar prescription_item.")
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("Error al comitear las transacciones: %w", err)
	}

	return prescriptionID, nil
}

func (r *postgreRepo) Read() ([]*PrescriptionBase, error) {
	ctx := context.Background()

	query := `
    SELECT
      p.prescription_id,
      (pt.first_name || ' ' || pt.last_name) AS patient_name,
      pt.dni AS patient_dni,
      p.issued_at,
      i.item_id,
      i.medication,
      i.dosage,
      i.duration_days,
      i.administration_route
    FROM prescription p
    JOIN patient pt
      ON pt.patient_id = p.patient_id
    JOIN prescription_item i
      ON i.prescription_id = p.prescription_id
    ORDER BY p.prescription_id, i.item_id;
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar consulta: %w", err)
	}
	defer rows.Close()

	var results []*PrescriptionBase
	var current *PrescriptionBase
	var lastID int

	for rows.Next() {
		var presID int
		var patientName, patientDNI string
		var issuedAt time.Time
		var item PrescriptionItemBase

		if err := rows.Scan(
			&presID,
			&patientName,
			&patientDNI,
			&issuedAt,
			&item.ItemID,
			&item.Medication,
			&item.Dosage,
			&item.DurationDays,
			&item.AdministrationRoute,
		); err != nil {
			return nil, fmt.Errorf("error al hacer Scan de filas: %w", err)
		}

		if current == nil || presID != lastID {
			current = &PrescriptionBase{
				PrescriptionID: presID,
				PatientName:    patientName,
				PatientDNI:     patientDNI,
				IssuedAt:       issuedAt,
				Items:          []PrescriptionItemBase{},
			}
			results = append(results, current)
			lastID = presID
		}

		current.Items = append(current.Items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error en la iteración de filas: %w", err)
	}

	return results, nil
}
