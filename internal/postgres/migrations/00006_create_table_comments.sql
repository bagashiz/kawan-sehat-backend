-- +goose Up
CREATE TABLE IF NOT EXISTS "comments" (
  "id" UUID PRIMARY KEY,
  "post_id" UUID NOT NULL,
  "account_id" UUID NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);
ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;
ALTER TABLE "comments" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS "comments";
