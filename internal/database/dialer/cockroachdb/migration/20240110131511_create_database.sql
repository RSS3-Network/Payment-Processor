-- +goose Up
-- +goose StatementBegin

-- public

-- public.checkpoints
CREATE TABLE "checkpoints"
(
    "chain_id"     bigint      NOT NULL,
    "block_number" bigint      NOT NULL,
    "block_hash"   text        NOT NULL,
    "created_at"   timestamptz NOT NULL DEFAULT now(),
    "updated_at"   timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT "pk_checkpoints" PRIMARY KEY ("chain_id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "checkpoints";
-- +goose StatementEnd
