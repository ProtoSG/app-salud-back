-- +goose Up
-- +goose StatementBegin
CREATE TABLE patient (
    patient_id        SERIAL       PRIMARY KEY,
    first_name        VARCHAR(100) NOT NULL,
    last_name         VARCHAR(100) NOT NULL,
    dni               VARCHAR(20)  NOT NULL UNIQUE,
    birth_date        DATE,
    gender            VARCHAR(20),
    address           TEXT,
    phone             VARCHAR(20),
    email             VARCHAR(150),
    photo_url         TEXT,
    registered_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS patient;
-- +goose StatementEnd
