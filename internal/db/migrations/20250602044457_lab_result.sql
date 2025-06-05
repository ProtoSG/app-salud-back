-- +goose Up
-- +goose StatementBegin
CREATE TABLE lab_result (
    lab_result_id    SERIAL       PRIMARY KEY,
    patient_id       INTEGER      NOT NULL,
    doctor_id        INTEGER,
    sample_date      DATE         NOT NULL,
    test_type        VARCHAR(100) NOT NULL,
    result_value     TEXT         NOT NULL,
    observations     TEXT,
    recorded_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES users(user_id)
);
INSERT INTO lab_result (
  patient_id, doctor_id, sample_date, test_type, result_value, observations, recorded_at
) VALUES
  (1, 2, '2023-01-09', 'Hemograma',   'Hb: 14 g/dL; Leucocitos: 7,000/µL', 'Todo dentro de rangos normales', NOW()),
  (2, 2, '2023-03-21', 'Perfil Lipídico', 'Colesterol total: 210 mg/dL',   'Ligeramente elevado', NOW()),
  (3, 2, '2023-05-07', 'Prueba de aliento', 'Negativo para H. pylori',       'Sin infección', NOW());


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lab_result;
-- +goose StatementEnd
