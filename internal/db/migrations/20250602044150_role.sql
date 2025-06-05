-- +goose Up
-- +goose StatementBegin
CREATE TABLE role (
    role_id      SERIAL       PRIMARY KEY,
    name         VARCHAR(50)  NOT NULL UNIQUE,
    description  TEXT
);

INSERT INTO role (name, description) VALUES
  ('ADMINISTRADOR', 'Rol de administrador'),
  ('DOCTOR',    'Rol de doctor'),
  ('ENFERMERO', 'Rol de enfermero');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role;
-- +goose StatementEnd
