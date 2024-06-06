-- +goose Up
-- modify "br_collected" table
ALTER TABLE "br_collected" ALTER COLUMN "index" SET NOT NULL;
-- modify "br_deposited" table
ALTER TABLE "br_deposited" ALTER COLUMN "index" SET NOT NULL;
-- modify "br_withdrawn" table
ALTER TABLE "br_withdrawn" ALTER COLUMN "index" SET NOT NULL;

-- +goose Down
-- reverse: modify "br_withdrawn" table
ALTER TABLE "br_withdrawn" ALTER COLUMN "index" DROP NOT NULL;
-- reverse: modify "br_deposited" table
ALTER TABLE "br_deposited" ALTER COLUMN "index" DROP NOT NULL;
-- reverse: modify "br_collected" table
ALTER TABLE "br_collected" ALTER COLUMN "index" DROP NOT NULL;
