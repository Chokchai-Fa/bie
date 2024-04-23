CREATE TABLE "User" (
  "email" varchar PRIMARY KEY,
  "hashed_password" varchar,
  "full_name" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "active" boolean,
  "authen_type" varchar
);

create table "Sessions" (
    "id" uuid primary key,
    "email" varchar not null,
    "refresh_token" varchar not null,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar not null,
    "is_blocked" boolean not null default false,
    "expired_at" timestamptz not null,
    "created_at" timestamptz not null default (now())
);

alter table "Sessions" add foreign key ("email") references "User" ("email");

CREATE TABLE "User_Activation"(
  "id" serial PRIMARY KEY,
  "email" varchar ,
  "activation_token" varchar NOT NULL
);

CREATE TABLE "User_Detail"(
  "user_id" serial PRIMARY KEY,
  "email" varchar,
  "subscribe_due_date" timestamptz,
  "payment_type" varchar,
  "omise_cust_id" varchar
);

ALTER TABLE "User_Activation" ADD FOREIGN KEY ("email") REFERENCES "User"("email");
ALTER TABLE "User_Detail" ADD FOREIGN KEY ("email") REFERENCES "User" ("email");

CREATE OR REPLACE VIEW "User_Roles_View" AS
SELECT
  user_id,
  email,
  CASE
    WHEN now() < subscribe_due_date THEN 'premium'
    ELSE 'normal'
  END AS user_role
FROM "User_Detail";
