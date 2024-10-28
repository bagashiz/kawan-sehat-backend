-- +goose Up
CREATE TABLE IF NOT EXISTS "replies" (
  "id" UUID PRIMARY KEY,
  "comment_id" UUID NOT NULL,
  "account_id" UUID NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);
ALTER TABLE "replies" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id") ON DELETE CASCADE;
ALTER TABLE "replies" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS "replies";
