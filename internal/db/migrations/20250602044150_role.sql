-- +goose Up
-- +goose StatementBegin
CREATE TABLE role (
    role_id      SERIAL       PRIMARY KEY,
    name         VARCHAR(50)  NOT NULL UNIQUE,
    description  TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rol;
-- +goose StatementEnd
