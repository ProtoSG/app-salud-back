package usuario

import (
	"database/sql"
	"fmt"

	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Repository interface {
	FindUser(email string) (bool, error)
	// ReadAppointmentToday(doctor_id int) error
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{
		db: db,
	}
}

func (r *postgreRepo) FindUser(email string) (bool, error) {
	const q = `
		SELECT 
			u.email
		FROM users u
		WHERE u.is_active = TRUE
		AND u.email = $1;
	`

	var emailUser string
	err := r.db.QueryRow(q, email).Scan(
		emailUser,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, utils.NewEntityNotFound(email, "email", "Usuario")
		}
		return false, fmt.Errorf("Error en la db: %w", err)
	}

	return true, nil
}
