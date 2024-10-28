-- +goose Up
CREATE TABLE IF NOT EXISTS "bookmarks" (
  "account_id" UUID NOT NULL,
  "post_id" UUID NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  PRIMARY KEY ("account_id", "post_id")
);
ALTER TABLE "bookmarks" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
ALTER TABLE "bookmarks" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS "bookmarks";
