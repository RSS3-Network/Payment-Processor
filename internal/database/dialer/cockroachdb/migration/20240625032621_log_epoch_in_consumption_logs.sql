-- +goose Up
-- modify "consumption_log" table
ALTER TABLE "consumption_log" ADD COLUMN "epoch" bigint NULL;
-- create index "idx_consumption_log_epoch" to table: "consumption_log"
CREATE INDEX "idx_consumption_log_epoch" ON "consumption_log" ("epoch");

-- +goose Down
-- reverse: create index "idx_consumption_log_epoch" to table: "consumption_log"
DROP INDEX "idx_consumption_log_epoch";
-- reverse: modify "consumption_log" table
ALTER TABLE "consumption_log" DROP COLUMN "epoch";
