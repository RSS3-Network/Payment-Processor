-- +goose Up
-- modify "br_collected" table
ALTER TABLE "br_collected" DROP CONSTRAINT "br_collected_pkey", ADD PRIMARY KEY ("tx_hash", "index");
-- modify "br_deposited" table
ALTER TABLE "br_deposited" DROP CONSTRAINT "br_deposited_pkey", ADD PRIMARY KEY ("tx_hash", "index");
-- modify "br_withdrawn" table
ALTER TABLE "br_withdrawn" DROP CONSTRAINT "br_withdrawn_pkey", ADD PRIMARY KEY ("tx_hash", "index");

-- +goose Down
-- reverse: modify "br_withdrawn" table
ALTER TABLE "br_withdrawn" DROP CONSTRAINT "br_withdrawn_pkey", ADD PRIMARY KEY ("tx_hash");
-- reverse: modify "br_deposited" table
ALTER TABLE "br_deposited" DROP CONSTRAINT "br_deposited_pkey", ADD PRIMARY KEY ("tx_hash");
-- reverse: modify "br_collected" table
ALTER TABLE "br_collected" DROP CONSTRAINT "br_collected_pkey", ADD PRIMARY KEY ("tx_hash");
