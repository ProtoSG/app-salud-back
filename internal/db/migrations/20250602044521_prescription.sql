-- +goose Up
-- +goose StatementBegin
CREATE TABLE prescription (
    prescription_id       SERIAL       PRIMARY KEY,
    patient_id            INTEGER      NOT NULL,
    doctor_id             INTEGER      NOT NULL,
    issued_at             TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    electronic_signature  TEXT,
    observations          TEXT,
    status                 VARCHAR(20)  NOT NULL DEFAULT 'ACTIVE',
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS prescription;
-- +goose StatementEnd
