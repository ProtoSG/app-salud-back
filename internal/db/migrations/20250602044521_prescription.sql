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

INSERT INTO prescription (
   patient_id, doctor_id, issued_at, electronic_signature, observations, status
) VALUES
  ( 1, 2, '2023-01-10 14:00:00', 'Dr. Juan Pérez', 'Revisar síntomas en una semana.', 'ACTIVE'),
  ( 2, 2, '2023-03-22 10:15:00', 'Dr. Juan Pérez', 'Seguir plan de dieta y medicamentos.', 'ACTIVE'),
  ( 3, 2, '2023-05-08 17:30:00', 'Dr. Juan Pérez', 'Controlar acidez y dolor.', 'COMPLETED');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS prescription;
-- +goose StatementEnd
