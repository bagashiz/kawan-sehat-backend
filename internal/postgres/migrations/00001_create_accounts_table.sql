-- +goose Up
CREATE TYPE ACCOUNT_ROLE AS ENUM ('PATIENT', 'EXPERT', 'ADMIN');
CREATE TYPE ACCOUNT_AVATAR AS ENUM ('NONE', 'OLD_FEMALE', 'OLD_MALE', 'YOUNG_FEMALE', 'YOUNG_MALE');
CREATE TYPE ACCOUNT_GENDER AS ENUM ('FEMALE', 'MALE', 'UNSPECIFIED');
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY,
    full_name VARCHAR(255),
    nik VARCHAR(16) UNIQUE,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    gender account_gender NOT NULL DEFAULT 'UNSPECIFIED',
    role account_role NOT NULL DEFAULT 'PATIENT',
    avatar account_avatar NOT NULL DEFAULT 'NONE',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
INSERT INTO accounts (
  id, full_name, nik, username,
  email, password, gender, role, 
  avatar, created_at, updated_at
) VALUES (
  '0192d7f8-46b7-7d44-ab16-128a6d32b4de',
  'Admin Name',
  '1234567890123456',
  'admin',
  'admin@example.com',
  '$2a$10$1mrzUWqCoBIGk/oxONkggOzMjZeGweEa7bqERs2BFLDq.dFBN0Jm.',
  'UNSPECIFIED',
  'ADMIN',
  'NONE',
  now(),
  now()
) ON CONFLICT (id) DO NOTHING;

-- +goose Down
DROP TABLE IF EXISTS accounts;
DROP TYPE IF EXISTS account_role;
DROP TYPE IF EXISTS account_avatar;
DROP TYPE IF EXISTS account_gender;
