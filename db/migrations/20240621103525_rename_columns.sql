-- Rename a column from "createdAt" to "created_at"
ALTER TABLE "public"."connections" RENAME COLUMN "createdAt" TO "created_at";
-- Rename a column from "updatedAt" to "updated_at"
ALTER TABLE "public"."connections" RENAME COLUMN "updatedAt" TO "updated_at";
-- Rename a column from "passwordHash" to "password_hash"
ALTER TABLE "public"."users" RENAME COLUMN "passwordHash" TO "password_hash";
-- Rename a column from "createdAt" to "created_at"
ALTER TABLE "public"."users" RENAME COLUMN "createdAt" TO "created_at";
-- Rename a column from "updatedAt" to "updated_at"
ALTER TABLE "public"."users" RENAME COLUMN "updatedAt" TO "updated_at";
-- Drop index "sessions_lastactivity_key" from table: "sessions"
DROP INDEX "public"."sessions_lastactivity_key";
-- Drop index "sessions_userid_key" from table: "sessions"
DROP INDEX "public"."sessions_userid_key";
-- Modify "sessions" table
ALTER TABLE "public"."sessions" DROP CONSTRAINT "sessions_userId_fkey", ADD
 CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Create index "sessions_lastactivity_key" to table: "sessions"
CREATE INDEX "sessions_lastactivity_key" ON "public"."sessions" ("last_activity");
-- Create index "sessions_userid_key" to table: "sessions"
CREATE INDEX "sessions_userid_key" ON "public"."sessions" ("user_id");
-- Rename a column from "userId" to "user_id"
ALTER TABLE "public"."sessions" RENAME COLUMN "userId" TO "user_id";
-- Rename a column from "ipAddress" to "ip_address"
ALTER TABLE "public"."sessions" RENAME COLUMN "ipAddress" TO "ip_address";
-- Rename a column from "userAgent" to "user_agent"
ALTER TABLE "public"."sessions" RENAME COLUMN "userAgent" TO "user_agent";
-- Rename a column from "lastActivity" to "last_activity"
ALTER TABLE "public"."sessions" RENAME COLUMN "lastActivity" TO "last_activity";
