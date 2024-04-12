-- +goose Up
-- create "billing_record_bases" table
CREATE TABLE "billing_record_bases" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "tx_hash" bytea NOT NULL,
  "index" bigint NULL,
  "chain_id" bigint NULL,
  "block_timestamp" timestamptz NULL,
  "block_number" bigint NULL,
  "user" bytea NULL,
  "amount" text NULL,
  PRIMARY KEY ("tx_hash")
);
-- create index "idx_billing_record_bases_block_number" to table: "billing_record_bases"
CREATE INDEX "idx_billing_record_bases_block_number" ON "billing_record_bases" ("block_number");
-- create index "idx_billing_record_bases_block_timestamp" to table: "billing_record_bases"
CREATE INDEX "idx_billing_record_bases_block_timestamp" ON "billing_record_bases" ("block_timestamp");
-- create index "idx_billing_record_bases_chain_id" to table: "billing_record_bases"
CREATE INDEX "idx_billing_record_bases_chain_id" ON "billing_record_bases" ("chain_id");
-- create "checkpoints" table
CREATE TABLE "checkpoints" (
  "chain_id" bigint NOT NULL DEFAULT unique_rowid(),
  "block_number" bigint NULL,
  "block_hash" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("chain_id")
);
-- create "account" table
CREATE TABLE "account" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "address" bytea NOT NULL,
  "ru_limit" bigint NULL,
  "is_paused" boolean NULL,
  "billing_rate" numeric NULL,
  PRIMARY KEY ("address")
);
-- create index "idx_account_deleted_at" to table: "account"
CREATE INDEX "idx_account_deleted_at" ON "account" ("deleted_at");
-- create "key" table
CREATE TABLE "key" (
  "id" bigint NOT NULL DEFAULT unique_rowid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "key" text NULL,
  "ru_used_total" bigint NULL,
  "ru_used_current" bigint NULL,
  "api_calls_total" bigint NULL,
  "api_calls_current" bigint NULL,
  "name" text NULL,
  "account_address" bytea NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_key_account" FOREIGN KEY ("account_address") REFERENCES "account" ("address") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_key_account_address" to table: "key"
CREATE INDEX "idx_key_account_address" ON "key" ("account_address");
-- create index "idx_key_deleted_at" to table: "key"
CREATE INDEX "idx_key_deleted_at" ON "key" ("deleted_at");
-- create index "idx_key_key" to table: "key"
CREATE UNIQUE INDEX "idx_key_key" ON "key" ("key");
-- create index "idx_key_ru_used_current" to table: "key"
CREATE INDEX "idx_key_ru_used_current" ON "key" ("ru_used_current");
-- create "consumption_log" table
CREATE TABLE "consumption_log" (
  "id" bigint NOT NULL DEFAULT unique_rowid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "consumption_date" timestamptz NULL,
  "ru_used" bigint NULL,
  "api_calls" bigint NULL,
  "key_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_consumption_log_key" FOREIGN KEY ("key_id") REFERENCES "key" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_consumption_log_consumption_date" to table: "consumption_log"
CREATE INDEX "idx_consumption_log_consumption_date" ON "consumption_log" ("consumption_date");
-- create index "idx_consumption_log_key_id" to table: "consumption_log"
CREATE INDEX "idx_consumption_log_key_id" ON "consumption_log" ("key_id");
-- create "pending_withdraw_request" table
CREATE TABLE "pending_withdraw_request" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "amount" numeric NULL,
  "account_address" bytea NOT NULL,
  PRIMARY KEY ("account_address"),
  CONSTRAINT "fk_pending_withdraw_request_account" FOREIGN KEY ("account_address") REFERENCES "account" ("address") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +goose Down
-- reverse: create "pending_withdraw_request" table
DROP TABLE "pending_withdraw_request";
-- reverse: create index "idx_consumption_log_key_id" to table: "consumption_log"
DROP INDEX "idx_consumption_log_key_id";
-- reverse: create index "idx_consumption_log_consumption_date" to table: "consumption_log"
DROP INDEX "idx_consumption_log_consumption_date";
-- reverse: create "consumption_log" table
DROP TABLE "consumption_log";
-- reverse: create index "idx_key_ru_used_current" to table: "key"
DROP INDEX "idx_key_ru_used_current";
-- reverse: create index "idx_key_key" to table: "key"
DROP INDEX "idx_key_key";
-- reverse: create index "idx_key_deleted_at" to table: "key"
DROP INDEX "idx_key_deleted_at";
-- reverse: create index "idx_key_account_address" to table: "key"
DROP INDEX "idx_key_account_address";
-- reverse: create "key" table
DROP TABLE "key";
-- reverse: create index "idx_account_deleted_at" to table: "account"
DROP INDEX "idx_account_deleted_at";
-- reverse: create "account" table
DROP TABLE "account";
-- reverse: create "checkpoints" table
DROP TABLE "checkpoints";
-- reverse: create index "idx_billing_record_bases_chain_id" to table: "billing_record_bases"
DROP INDEX "idx_billing_record_bases_chain_id";
-- reverse: create index "idx_billing_record_bases_block_timestamp" to table: "billing_record_bases"
DROP INDEX "idx_billing_record_bases_block_timestamp";
-- reverse: create index "idx_billing_record_bases_block_number" to table: "billing_record_bases"
DROP INDEX "idx_billing_record_bases_block_number";
-- reverse: create "billing_record_bases" table
DROP TABLE "billing_record_bases";
