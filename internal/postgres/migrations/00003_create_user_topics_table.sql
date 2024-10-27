-- +goose Up
CREATE TABLE IF NOT EXISTS "account_topics" (
  "account_id" UUID NOT NULL,
  "topic_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  PRIMARY KEY ("account_id", "topic_id")
);
ALTER TABLE "account_topics" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
ALTER TABLE "account_topics" ADD FOREIGN KEY ("topic_id") REFERENCES "topics" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS "account_topics";
