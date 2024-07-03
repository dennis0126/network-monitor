-- Modify "sessions" table
ALTER TABLE "public"."sessions" ALTER COLUMN "last_activity" TYPE timestamptz;
-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz;
-- Modify "connections" table
ALTER TABLE "public"."connections" DROP CONSTRAINT "owner_id", ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz, ADD
 CONSTRAINT "connections_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
