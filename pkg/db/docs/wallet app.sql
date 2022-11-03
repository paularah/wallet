CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar,
  "ip" varchar,
  "location" varchar,
  "is_valid" boolean NOT NULL DEFAULT true,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "fundings" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "wallet_id" bigserial NOT NULL,
  "amount" bigint NOT NULL,
  "success" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "owner" bigint NOT NULL,
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

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "sessions" ("user_id");

CREATE INDEX ON "sessions" ("id");

CREATE INDEX ON "fundings" ("id");

CREATE INDEX ON "fundings" ("user_id");

CREATE INDEX ON "fundings" ("wallet_id");

CREATE INDEX ON "wallets" ("owner");

CREATE UNIQUE INDEX ON "wallets" ("owner", "currency");

CREATE INDEX ON "entries" ("wallet_id");

CREATE INDEX ON "transfers" ("sender_wallet_id");

CREATE INDEX ON "transfers" ("receiver_wallet_id");

CREATE INDEX ON "transfers" ("sender_wallet_id", "receiver_wallet_id");

COMMENT ON COLUMN "entries"."amount" IS 'negative values indicate deductions';

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "fundings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "fundings" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "wallets" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("sender_wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("receiver_wallet_id") REFERENCES "wallets" ("id");
