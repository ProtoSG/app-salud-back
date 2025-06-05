package auth

import (
	"database/sql"
	"fmt"

	"github.com/ProtoSG/app-salud-back/internal/features/user"
	"github.com/ProtoSG/app-salud-back/internal/utils"
)

type Repository interface {
	CreateUser(user *user.User) (int64, error)
	FindUserByEmail(email string) (*UserAuth, error)
}

type postgreRepo struct {
	db *sql.DB
}

func NewPostgreRepo(db *sql.DB) Repository {
	return &postgreRepo{db}
}

func (r *postgreRepo) CreateUser(user *user.User) (int64, error) {
	const q = `
		INSERT INTO users (role_id, first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id
	`

	var id int64

	err := r.db.QueryRow(
		q,
		user.RoleID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repo Create: %w", err)
	}

	return id, nil
}

func (r *postgreRepo) FindUserByEmail(email string) (*UserAuth, error) {
	const q = `
		SELECT 
			u.user_id, 
			u.email,
			u.password,
			r.name
		FROM users u
		INNER JOIN role r ON r.role_id = u.role_id
		WHERE u.is_active = TRUE
		AND u.email = $1;
	`

	userAuth := &UserAuth{}
	err := r.db.QueryRow(q, email).Scan(
		&userAuth.UserID,
		&userAuth.Email,
		&userAuth.Password,
		&userAuth.RolName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NewEntityNotFound(email, "email", "Usuario")
		}
		return nil, fmt.Errorf("Error en la db: %w", err)
	}

	return userAuth, nil
}
