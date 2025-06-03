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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS prescription_item;
-- +goose StatementEnd
