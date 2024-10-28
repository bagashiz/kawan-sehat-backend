-- +goose Up
CREATE TABLE IF NOT EXISTS "posts" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "topic_id" UUID NOT NULL,
  "title" VARCHAR(255) NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);
ALTER TABLE "posts" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
ALTER TABLE "posts" ADD FOREIGN KEY ("topic_id") REFERENCES "topics" ("id") ON DELETE CASCADE;


-- +goose Down
DROP TABLE IF EXISTS "posts";
