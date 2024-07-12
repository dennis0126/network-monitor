-- Create enum type "connection_type"
CREATE TYPE "public"."connection_type" AS ENUM ('ping', 'http');
-- Modify "connections" table
ALTER TABLE "public"."connections" ADD COLUMN "type" "public"."connection_type" NOT NULL DEFAULT 'ping', ADD COLUMN "interval" integer NOT NULL DEFAULT 1;
-- Create "connections_http" table
CREATE TABLE "public"."connections_http" (
  "id" character varying(36) NOT NULL,
  "type" "public"."connection_type" NOT NULL DEFAULT 'http',
  "destination" character varying(100) NOT NULL,
  CONSTRAINT "connections_http_id_fkey" FOREIGN KEY ("id") REFERENCES "public"."connections" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "connections_ping" table
CREATE TABLE "public"."connections_ping" (
  "id" character varying(36) NOT NULL,
  "type" "public"."connection_type" NOT NULL DEFAULT 'ping',
  "destination" character varying(100) NOT NULL,
  CONSTRAINT "connections_ping_id_fkey" FOREIGN KEY ("id") REFERENCES "public"."connections" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
