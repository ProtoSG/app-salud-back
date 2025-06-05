package prescription

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(pres *Prescription) (int64, error)
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
