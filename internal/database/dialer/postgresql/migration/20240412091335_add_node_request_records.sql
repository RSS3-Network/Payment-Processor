-- +goose Up
-- create "node_request_record" table
CREATE TABLE "node_request_record" (
  "epoch" bigint NOT NULL,
  "node_address" bytea NOT NULL,
  "request_counts" bigint NULL,
  "request_rewards" text NULL,
  PRIMARY KEY ("epoch", "node_address")
);

-- +goose Down
-- reverse: create "node_request_record" table
DROP TABLE "node_request_record";
