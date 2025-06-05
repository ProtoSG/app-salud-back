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

INSERT INTO medical_appointment (
  patient_id, doctor_id, administrative_id, appointment_time, duration_minutes,
  reason, status, created_at, updated_at, additional_notes
) VALUES
  (1, 2, 1, '2023-09-01 09:30:00', 30, 'Revisión postvacuna COVID-19', 'COMPLETED', NOW(), NOW(), 'Paciente estable.'),
  (2, 2, 1, '2023-09-05 11:00:00', 45, 'Control de hipertensión',      'SCHEDULED', NOW(), NOW(), 'Medir presión.'),
  (3, 2, 1, '2023-09-10 16:00:00', 30, 'Consulta por dolor estomacal', 'CANCELLED', NOW(), NOW(), 'El paciente pospuso.' );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS medical_appointment;
-- +goose StatementEnd
