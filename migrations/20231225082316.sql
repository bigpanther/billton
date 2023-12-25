-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(255) NOT NULL,
  "username" character varying(255) NOT NULL,
  "email" character varying(255) NOT NULL,
  "role" character varying(50) NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create "warranties" table
CREATE TABLE "public"."warranties" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "transaction_time" timestamptz NULL,
  "expiry_time" timestamptz NULL,
  "brand_name" character varying(255) NOT NULL,
  "store_name" character varying(255) NOT NULL,
  "amount" bigint NULL,
  "user_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_warranties" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_warranties_deleted_at" to table: "warranties"
CREATE INDEX "idx_warranties_deleted_at" ON "public"."warranties" ("deleted_at");
-- Create index "idx_warranties_user_id" to table: "warranties"
CREATE INDEX "idx_warranties_user_id" ON "public"."warranties" ("user_id");
