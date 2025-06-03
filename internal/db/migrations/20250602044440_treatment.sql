-- +goose Up
-- +goose StatementBegin
CREATE TABLE treatment (
    treatment_id     SERIAL       PRIMARY KEY,
    patient_id       INTEGER      NOT NULL,
    doctor_id        INTEGER      NOT NULL,
    start_date       DATE         NOT NULL,
    end_date         DATE,
    description      TEXT         NOT NULL,
    observations     TEXT,
    recorded_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS treatment;
-- +goose StatementEnd
