CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "owner" bigint UNIQUE NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "wallet_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "sender_wallet_id" bigint NOT NULL,
  "receiver_wallet_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "wallets" ("owner");

CREATE INDEX ON "entries" ("wallet_id");

CREATE INDEX ON "transfers" ("sender_wallet_id");

CREATE INDEX ON "transfers" ("receiver_wallet_id");

CREATE INDEX ON "transfers" ("sender_wallet_id", "receiver_wallet_id");

COMMENT ON COLUMN "entries"."amount" IS 'negative values indicate deductions';

ALTER TABLE "entries" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("sender_wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("receiver_wallet_id") REFERENCES "wallets" ("id");
