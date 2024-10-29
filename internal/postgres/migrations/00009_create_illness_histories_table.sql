-- +goose Up
CREATE TABLE IF NOT EXISTS "illness_histories" (
    account_id UUID NOT NULL,
    illness VARCHAR(255) NOT NULL,
    date DATE NOT NULL
);
ALTER TABLE "illness_histories" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS "illness_histories";
