-- +goose Up
-- +goose StatementBegin
CREATE TABLE vaccine (
    vaccine_id       SERIAL       PRIMARY KEY,
    patient_id       INTEGER      NOT NULL,
    vaccine_type     VARCHAR(100) NOT NULL,
    administered_on  DATE         NOT NULL,
    dose             VARCHAR(50),
    observations     TEXT,
    recorded_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vaccine;
-- +goose StatementEnd
