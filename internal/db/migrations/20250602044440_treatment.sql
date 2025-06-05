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

INSERT INTO treatment (
   patient_id, doctor_id, start_date, end_date, description, observations, recorded_at
) VALUES
  ( 1, 2, '2023-01-11', '2023-01-18', 'Reposo domiciliario', 'Recetar paracetamol si fiebre >38°C', NOW()),
  ( 2, 2, '2023-03-23', '2023-06-23', 'Tratamiento para hipertensión', 'Monitorear presión cada 2 semanas', NOW()),
  ( 3, 2, '2023-05-09', '2023-05-23', 'Tratamiento gastrológico', 'Evitar comidas picantes', NOW());


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS treatment;
-- +goose StatementEnd
