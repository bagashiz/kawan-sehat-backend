-- +goose Up
CREATE TYPE user_role AS ENUM ('patient', 'expert', 'admin');
CREATE TYPE user_avatar AS ENUM ('old_female', 'old_male', 'young_female', 'young_male');
CREATE TYPE user_gender AS ENUM ('female', 'male', 'unspecified');
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nik VARCHAR(16) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    gender user_gender NOT NULL DEFAULT 'unspecified',
    role user_role NOT NULL DEFAULT 'patient',
    avatar user_avatar,
    illness_history TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;
DROP TYPE IF EXISTS user_avatar;
DROP TYPE IF EXISTS user_gender;
