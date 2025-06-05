-- +goose Up
-- +goose StatementBegin
CREATE TABLE patient (
    patient_id        SERIAL       PRIMARY KEY,
    first_name        VARCHAR(100) NOT NULL,
    last_name         VARCHAR(100) NOT NULL,
    dni               VARCHAR(20)  NOT NULL UNIQUE,
    birth_date        DATE,
    gender            VARCHAR(20),
    address           TEXT,
    phone             VARCHAR(20),
    email             VARCHAR(150),
    photo_url         TEXT,
    registered_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

INSERT INTO patient (
  first_name, last_name, dni, birth_date, gender,
  address, phone, email, photo_url, registered_at, is_deleted
) VALUES
  ('María',    'González',   '12345678A', '1985-07-12', 'F', 'C/ Falsa, 123, Madrid',    '600123456', 'maria@example.com',    NULL, NOW(), FALSE),
  ('Carlos',   'Ramírez',    '87654321B', '1990-03-05', 'M', 'Av. Siempre Viva, 45, Sevilla', '600654321', 'carlos@example.com',   NULL, NOW(), FALSE),
  ('Luisa',    'Fernández',  '23456789C', '1978-11-23', 'F', 'Plaza Mayor, 5, Valencia',     '600789012', 'luisa@example.com',    NULL, NOW(), FALSE);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS patient;
-- +goose StatementEnd
