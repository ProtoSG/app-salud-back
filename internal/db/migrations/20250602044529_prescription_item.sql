-- +goose Up
-- +goose StatementBegin
CREATE TABLE prescription_item (
    item_id            SERIAL       PRIMARY KEY,
    prescription_id    INTEGER      NOT NULL,
    medication         VARCHAR(150) NOT NULL,
    dosage             VARCHAR(50)  NOT NULL,
    frequency          VARCHAR(100) NOT NULL,
    duration_days      INTEGER      NOT NULL,
    administration_route VARCHAR(50) NOT NULL,
    observations       TEXT,
    FOREIGN KEY(prescription_id) REFERENCES prescription(prescription_id)
);
INSERT INTO prescription_item (
  prescription_id, medication, dosage, frequency, duration_days, administration_route, observations
) VALUES
  (1, 'Paracetamol 500 mg', '500 mg',        'Cada 8 horas', 5, 'Oral', 'Tomar con comida'),
  (2, 'Enalapril 10 mg',    '10 mg',         'Una vez al día', 90, 'Oral', 'Tomar en ayunas'),
  (2, 'Amlodipino 5 mg',    '5 mg',          'Una vez al día', 90, 'Oral', 'Si presión > 140/90'),
  (3, 'Omeprazol 20 mg',    '20 mg',         'Cada 12 horas', 14, 'Oral', '20 minutos antes de comer');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS prescription_item;
-- +goose StatementEnd
