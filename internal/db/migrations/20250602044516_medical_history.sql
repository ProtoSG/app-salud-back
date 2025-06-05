-- +goose Up
-- +goose StatementBegin
CREATE TABLE medical_history (
    history_id       SERIAL       PRIMARY KEY,
    patient_id       INTEGER      NOT NULL,
    description      TEXT         NOT NULL,
    recorded_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id)
);
INSERT INTO medical_history (
  patient_id, description, recorded_at
) VALUES
  (1, 'Alergia a penicilina.', NOW()),
  (2, 'Historial de hipertensión arterial.', NOW()),
  (3, 'Operada de apendicitis en 2005.', NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS medical_history;
-- +goose StatementEnd
