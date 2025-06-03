-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id            SERIAL       PRIMARY KEY,
    role_id            INTEGER      NOT NULL,
    first_name         VARCHAR(100) NOT NULL,
    last_name          VARCHAR(100) NOT NULL,
    email              VARCHAR(150) NOT NULL UNIQUE,
    password           VARCHAR(255) NOT NULL,
    created_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active          BOOLEAN      NOT NULL DEFAULT TRUE,
    FOREIGN KEY (role_id) REFERENCES role(role_id)
);

CREATE INDEX idx_user_email ON users(email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
