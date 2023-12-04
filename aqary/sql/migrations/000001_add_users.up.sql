--liquibase formatted sql
-- changeset mobin:1

CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "name" TEXT NOT NULL,
  "phone_number" TEXT NOT NULL UNIQUE,
  "otp" TEXT NULL,
  "otp_expiration_time" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("phone_number");
--rollback DROP TABLE "users";


