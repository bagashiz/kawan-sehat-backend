-- +goose Up
CREATE TYPE account_role AS ENUM ('patient', 'expert', 'admin');
CREATE TYPE account_avatar AS ENUM ('old_female', 'old_male', 'young_female', 'young_male');
CREATE TYPE account_gender AS ENUM ('female', 'male', 'unspecified');
CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nik VARCHAR(16) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    gender account_gender NOT NULL DEFAULT 'unspecified',
    role account_role NOT NULL DEFAULT 'patient',
    avatar account_avatar,
    illness_history TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS accounts;
DROP TYPE IF EXISTS account_role;
DROP TYPE IF EXISTS account_avatar;
DROP TYPE IF EXISTS account_gender;
