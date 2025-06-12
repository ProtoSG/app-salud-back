package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	goose.AddMigrationContext(upUsersSeed, downUsersSeed)
}

func upUsersSeed(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	pwd := "test"
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	hashStr := string(hash)

	rows := []struct {
		roleID    int
		firstName string
		lastName  string
		email     string
	}{
		{1, "Admin", "User", "admin@test.com"},
		{2, "Doctor", "Who", "doctor@test.com"},
		{3, "Enfermero", "Salud", "enfermero@test.com"},
	}

	query := `
    INSERT INTO users (role_id, first_name, last_name, email, password) VALUES
      ($1,  $2,  $3,  $4,  $5),
      ($6,  $7,  $8,  $9,  $10),
      ($11, $12, $13, $14, $15);
    `

	args := []any{
		rows[0].roleID, rows[0].firstName, rows[0].lastName, rows[0].email, hashStr,
		rows[1].roleID, rows[1].firstName, rows[1].lastName, rows[1].email, hashStr,
		rows[2].roleID, rows[2].firstName, rows[2].lastName, rows[2].email, hashStr,
	}

	if _, err := tx.Exec(query, args...); err != nil {
		return fmt.Errorf("insert users seed: %w", err)
	}

	return nil
}

func downUsersSeed(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
