-- +goose Up
-- +goose StatementBegin
CREATE TABLE medical_history (
    history_id       SERIAL       PRIMARY KEY,
    patient_id       INTEGER      NOT NULL,
    description      TEXT         NOT NULL,
    recorded_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS medical_history;
-- +goose StatementEnd
