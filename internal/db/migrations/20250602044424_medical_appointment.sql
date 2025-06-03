-- +goose Up
-- +goose StatementBegin
CREATE TABLE medical_appointment (
    appointment_id      SERIAL       PRIMARY KEY,
    patient_id          INTEGER      NOT NULL,
    doctor_id           INTEGER      NOT NULL,
    administrative_id   INTEGER,
    appointment_time    TIMESTAMP    NOT NULL,
    duration_minutes    INTEGER      NOT NULL,
    reason              TEXT,
    status              VARCHAR(20)  NOT NULL DEFAULT 'SCHEDULED',
    created_at          TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP,
    additional_notes    TEXT,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES users(user_id),
    FOREIGN KEY (administrative_id) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS medical_appointment;
-- +goose StatementEnd
