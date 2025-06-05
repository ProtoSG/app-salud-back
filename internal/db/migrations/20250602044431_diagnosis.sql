-- +goose Up
-- +goose StatementBegin
CREATE TABLE diagnosis (
    diagnosis_id      SERIAL       PRIMARY KEY,
    patient_id        INTEGER      NOT NULL,
    doctor_id         INTEGER      NOT NULL,
    diagnosis_date    DATE         NOT NULL,
    description       TEXT         NOT NULL,
    observations      TEXT,
    recorded_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES users(user_id)
);
INSERT INTO diagnosis (
  patient_id, doctor_id, diagnosis_date, description, observations, recorded_at
) VALUES
  ( 1, 2, '2023-01-10', 'Gripe estacional',       'Tratamiento con reposo y fluidos',   NOW()),
  ( 2, 2, '2023-03-22', 'Hipertensión leve',      'Controlar presión y dieta baja en sodio', NOW()),
  ( 3, 2, '2023-05-08', 'Gastritis aguda',       'Prescribir omeprazol',                NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS diagnosis;
-- +goose StatementEnd
