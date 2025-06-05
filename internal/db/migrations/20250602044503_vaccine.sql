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
INSERT INTO vaccine (
  patient_id, vaccine_type, administered_on, dose, observations, recorded_at
) VALUES
  (1, 'COVID-19',      '2021-06-15', '1ª dosis',      'Sin efectos adversos', NOW()),
  (1, 'COVID-19',      '2021-07-15', '2ª dosis',      'Ligera fiebre', NOW()),
  (2, 'Influenza',     '2022-10-01', 'Única dosis',   'Vacuna anual', NOW()),
  (3, 'Hepatitis B',   '2020-02-20', '3ª dosis',      'Completo esquema', NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vaccine;
-- +goose StatementEnd
