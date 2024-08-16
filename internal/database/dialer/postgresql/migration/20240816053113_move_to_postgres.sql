-- +goose Up
-- create "br_collected" table
CREATE TABLE "public"."br_collected" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "tx_hash" bytea NOT NULL,
  "index" bigint NOT NULL,
  "chain_id" bigint NULL,
  "block_timestamp" timestamptz NULL,
  "block_number" bigint NULL,
  "user" bytea NULL,
  "amount" text NULL,
  PRIMARY KEY ("tx_hash", "index")
);
-- create index "idx_br_collected_block_number" to table: "br_collected"
CREATE INDEX "idx_br_collected_block_number" ON "public"."br_collected" ("block_number");
-- create index "idx_br_collected_block_timestamp" to table: "br_collected"
CREATE INDEX "idx_br_collected_block_timestamp" ON "public"."br_collected" ("block_timestamp");
-- create index "idx_br_collected_chain_id" to table: "br_collected"
CREATE INDEX "idx_br_collected_chain_id" ON "public"."br_collected" ("chain_id");
-- create "br_deposited" table
CREATE TABLE "public"."br_deposited" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "tx_hash" bytea NOT NULL,
  "index" bigint NOT NULL,
  "chain_id" bigint NULL,
  "block_timestamp" timestamptz NULL,
  "block_number" bigint NULL,
  "user" bytea NULL,
  "amount" text NULL,
  PRIMARY KEY ("tx_hash", "index")
);
-- create index "idx_br_deposited_block_number" to table: "br_deposited"
CREATE INDEX "idx_br_deposited_block_number" ON "public"."br_deposited" ("block_number");
-- create index "idx_br_deposited_block_timestamp" to table: "br_deposited"
CREATE INDEX "idx_br_deposited_block_timestamp" ON "public"."br_deposited" ("block_timestamp");
-- create index "idx_br_deposited_chain_id" to table: "br_deposited"
CREATE INDEX "idx_br_deposited_chain_id" ON "public"."br_deposited" ("chain_id");
-- create "br_withdrawn" table
CREATE TABLE "public"."br_withdrawn" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "tx_hash" bytea NOT NULL,
  "index" bigint NOT NULL,
  "chain_id" bigint NULL,
  "block_timestamp" timestamptz NULL,
  "block_number" bigint NULL,
  "user" bytea NULL,
  "amount" text NULL,
  "fee" text NULL,
  PRIMARY KEY ("tx_hash", "index")
);
-- create index "idx_br_withdrawn_block_number" to table: "br_withdrawn"
CREATE INDEX "idx_br_withdrawn_block_number" ON "public"."br_withdrawn" ("block_number");
-- create index "idx_br_withdrawn_block_timestamp" to table: "br_withdrawn"
CREATE INDEX "idx_br_withdrawn_block_timestamp" ON "public"."br_withdrawn" ("block_timestamp");
-- create index "idx_br_withdrawn_chain_id" to table: "br_withdrawn"
CREATE INDEX "idx_br_withdrawn_chain_id" ON "public"."br_withdrawn" ("chain_id");
-- create "checkpoints" table
CREATE TABLE "public"."checkpoints" (
  "chain_id" bigserial NOT NULL,
  "block_number" bigint NULL,
  "block_hash" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("chain_id")
);
-- create "node_request_record" table
CREATE TABLE "public"."node_request_record" (
  "epoch" bigint NOT NULL,
  "node_address" bytea NOT NULL,
  "request_counts" bigint NULL,
  "request_rewards" text NULL,
  PRIMARY KEY ("epoch", "node_address")
);
-- create "account" table
CREATE TABLE "public"."account" (
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
CREATE INDEX "idx_account_deleted_at" ON "public"."account" ("deleted_at");
-- create "key" table
CREATE TABLE "public"."key" (
  "id" bigserial NOT NULL,
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
  CONSTRAINT "fk_key_account" FOREIGN KEY ("account_address") REFERENCES "public"."account" ("address") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_key_account_address" to table: "key"
CREATE INDEX "idx_key_account_address" ON "public"."key" ("account_address");
-- create index "idx_key_deleted_at" to table: "key"
CREATE INDEX "idx_key_deleted_at" ON "public"."key" ("deleted_at");
-- create index "idx_key_key" to table: "key"
CREATE UNIQUE INDEX "idx_key_key" ON "public"."key" ("key");
-- create index "idx_key_ru_used_current" to table: "key"
CREATE INDEX "idx_key_ru_used_current" ON "public"."key" ("ru_used_current");
-- create "consumption_log" table
CREATE TABLE "public"."consumption_log" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "consumption_date" timestamptz NULL,
  "epoch" bigint NULL,
  "ru_used" bigint NULL,
  "api_calls" bigint NULL,
  "key_id" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_consumption_log_key" FOREIGN KEY ("key_id") REFERENCES "public"."key" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_consumption_log_consumption_date" to table: "consumption_log"
CREATE INDEX "idx_consumption_log_consumption_date" ON "public"."consumption_log" ("consumption_date");
-- create index "idx_consumption_log_epoch" to table: "consumption_log"
CREATE INDEX "idx_consumption_log_epoch" ON "public"."consumption_log" ("epoch");
-- create index "idx_consumption_log_key_id" to table: "consumption_log"
CREATE INDEX "idx_consumption_log_key_id" ON "public"."consumption_log" ("key_id");
-- create "pending_withdraw_request" table
CREATE TABLE "public"."pending_withdraw_request" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "amount" numeric NULL,
  "account_address" bytea NOT NULL,
  PRIMARY KEY ("account_address"),
  CONSTRAINT "fk_pending_withdraw_request_account" FOREIGN KEY ("account_address") REFERENCES "public"."account" ("address") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +goose Down
-- reverse: create "pending_withdraw_request" table
DROP TABLE "public"."pending_withdraw_request";
-- reverse: create index "idx_consumption_log_key_id" to table: "consumption_log"
DROP INDEX "public"."idx_consumption_log_key_id";
-- reverse: create index "idx_consumption_log_epoch" to table: "consumption_log"
DROP INDEX "public"."idx_consumption_log_epoch";
-- reverse: create index "idx_consumption_log_consumption_date" to table: "consumption_log"
DROP INDEX "public"."idx_consumption_log_consumption_date";
-- reverse: create "consumption_log" table
DROP TABLE "public"."consumption_log";
-- reverse: create index "idx_key_ru_used_current" to table: "key"
DROP INDEX "public"."idx_key_ru_used_current";
-- reverse: create index "idx_key_key" to table: "key"
DROP INDEX "public"."idx_key_key";
-- reverse: create index "idx_key_deleted_at" to table: "key"
DROP INDEX "public"."idx_key_deleted_at";
-- reverse: create index "idx_key_account_address" to table: "key"
DROP INDEX "public"."idx_key_account_address";
-- reverse: create "key" table
DROP TABLE "public"."key";
-- reverse: create index "idx_account_deleted_at" to table: "account"
DROP INDEX "public"."idx_account_deleted_at";
-- reverse: create "account" table
DROP TABLE "public"."account";
-- reverse: create "node_request_record" table
DROP TABLE "public"."node_request_record";
-- reverse: create "checkpoints" table
DROP TABLE "public"."checkpoints";
-- reverse: create index "idx_br_withdrawn_chain_id" to table: "br_withdrawn"
DROP INDEX "public"."idx_br_withdrawn_chain_id";
-- reverse: create index "idx_br_withdrawn_block_timestamp" to table: "br_withdrawn"
DROP INDEX "public"."idx_br_withdrawn_block_timestamp";
-- reverse: create index "idx_br_withdrawn_block_number" to table: "br_withdrawn"
DROP INDEX "public"."idx_br_withdrawn_block_number";
-- reverse: create "br_withdrawn" table
DROP TABLE "public"."br_withdrawn";
-- reverse: create index "idx_br_deposited_chain_id" to table: "br_deposited"
DROP INDEX "public"."idx_br_deposited_chain_id";
-- reverse: create index "idx_br_deposited_block_timestamp" to table: "br_deposited"
DROP INDEX "public"."idx_br_deposited_block_timestamp";
-- reverse: create index "idx_br_deposited_block_number" to table: "br_deposited"
DROP INDEX "public"."idx_br_deposited_block_number";
-- reverse: create "br_deposited" table
DROP TABLE "public"."br_deposited";
-- reverse: create index "idx_br_collected_chain_id" to table: "br_collected"
DROP INDEX "public"."idx_br_collected_chain_id";
-- reverse: create index "idx_br_collected_block_timestamp" to table: "br_collected"
DROP INDEX "public"."idx_br_collected_block_timestamp";
-- reverse: create index "idx_br_collected_block_number" to table: "br_collected"
DROP INDEX "public"."idx_br_collected_block_number";
-- reverse: create "br_collected" table
DROP TABLE "public"."br_collected";
