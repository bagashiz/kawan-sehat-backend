-- +goose Up
CREATE TABLE IF NOT EXISTS "votes" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "post_id" UUID,
  "comment_id" UUID,
  "reply_id" UUID,
  "value" SMALLINT NOT NULL CHECK ( value IN (1, -1) ),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);
CREATE UNIQUE INDEX ON "votes" ("account_id", "post_id", "comment_id", "reply_id");
ALTER TABLE "votes" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE;
ALTER TABLE "votes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;
ALTER TABLE "votes" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id") ON DELETE CASCADE;
ALTER TABLE "votes" ADD FOREIGN KEY ("reply_id") REFERENCES "replies" ("id") ON DELETE CASCADE;

-- +goose Down
DROP TABLE IF EXISTS "votes";
