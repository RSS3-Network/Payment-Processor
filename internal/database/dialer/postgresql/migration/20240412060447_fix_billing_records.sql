-- +goose Up
-- create "br_collected" table
CREATE TABLE "br_collected" (
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
-- create index "idx_br_collected_block_number" to table: "br_collected"
CREATE INDEX "idx_br_collected_block_number" ON "br_collected" ("block_number");
-- create index "idx_br_collected_block_timestamp" to table: "br_collected"
CREATE INDEX "idx_br_collected_block_timestamp" ON "br_collected" ("block_timestamp");
-- create index "idx_br_collected_chain_id" to table: "br_collected"
CREATE INDEX "idx_br_collected_chain_id" ON "br_collected" ("chain_id");
-- create "br_deposited" table
CREATE TABLE "br_deposited" (
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
-- create index "idx_br_deposited_block_number" to table: "br_deposited"
CREATE INDEX "idx_br_deposited_block_number" ON "br_deposited" ("block_number");
-- create index "idx_br_deposited_block_timestamp" to table: "br_deposited"
CREATE INDEX "idx_br_deposited_block_timestamp" ON "br_deposited" ("block_timestamp");
-- create index "idx_br_deposited_chain_id" to table: "br_deposited"
CREATE INDEX "idx_br_deposited_chain_id" ON "br_deposited" ("chain_id");
-- create "br_withdrawn" table
CREATE TABLE "br_withdrawn" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "tx_hash" bytea NOT NULL,
  "index" bigint NULL,
  "chain_id" bigint NULL,
  "block_timestamp" timestamptz NULL,
  "block_number" bigint NULL,
  "user" bytea NULL,
  "amount" text NULL,
  "fee" text NULL,
  PRIMARY KEY ("tx_hash")
);
-- create index "idx_br_withdrawn_block_number" to table: "br_withdrawn"
CREATE INDEX "idx_br_withdrawn_block_number" ON "br_withdrawn" ("block_number");
-- create index "idx_br_withdrawn_block_timestamp" to table: "br_withdrawn"
CREATE INDEX "idx_br_withdrawn_block_timestamp" ON "br_withdrawn" ("block_timestamp");
-- create index "idx_br_withdrawn_chain_id" to table: "br_withdrawn"
CREATE INDEX "idx_br_withdrawn_chain_id" ON "br_withdrawn" ("chain_id");
-- drop "billing_record_bases" table
DROP TABLE "billing_record_bases";

-- +goose Down
-- reverse: drop "billing_record_bases" table
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
CREATE INDEX "idx_billing_record_bases_block_number" ON "billing_record_bases" ("block_number");
CREATE INDEX "idx_billing_record_bases_block_timestamp" ON "billing_record_bases" ("block_timestamp");
CREATE INDEX "idx_billing_record_bases_chain_id" ON "billing_record_bases" ("chain_id");
-- reverse: create index "idx_br_withdrawn_chain_id" to table: "br_withdrawn"
DROP INDEX "idx_br_withdrawn_chain_id";
-- reverse: create index "idx_br_withdrawn_block_timestamp" to table: "br_withdrawn"
DROP INDEX "idx_br_withdrawn_block_timestamp";
-- reverse: create index "idx_br_withdrawn_block_number" to table: "br_withdrawn"
DROP INDEX "idx_br_withdrawn_block_number";
-- reverse: create "br_withdrawn" table
DROP TABLE "br_withdrawn";
-- reverse: create index "idx_br_deposited_chain_id" to table: "br_deposited"
DROP INDEX "idx_br_deposited_chain_id";
-- reverse: create index "idx_br_deposited_block_timestamp" to table: "br_deposited"
DROP INDEX "idx_br_deposited_block_timestamp";
-- reverse: create index "idx_br_deposited_block_number" to table: "br_deposited"
DROP INDEX "idx_br_deposited_block_number";
-- reverse: create "br_deposited" table
DROP TABLE "br_deposited";
-- reverse: create index "idx_br_collected_chain_id" to table: "br_collected"
DROP INDEX "idx_br_collected_chain_id";
-- reverse: create index "idx_br_collected_block_timestamp" to table: "br_collected"
DROP INDEX "idx_br_collected_block_timestamp";
-- reverse: create index "idx_br_collected_block_number" to table: "br_collected"
DROP INDEX "idx_br_collected_block_number";
-- reverse: create "br_collected" table
DROP TABLE "br_collected";
