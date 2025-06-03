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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lab_result;
-- +goose StatementEnd
