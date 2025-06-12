-- +goose Up
-- +goose StatementBegin
CREATE TABLE allergy (
    allergy_id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id)
);
INSERT INTO allergy (patient_id, name) VALUES
  (1, 'Penicillin'),
  (2, 'Peanuts'),
  (3, 'Latex');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS allergy;
-- +goose StatementEnd
