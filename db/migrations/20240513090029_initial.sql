-- Create "users" table
CREATE TABLE "public"."users" (
  "id" character varying(36) NOT NULL,
  "name" character varying(100) NOT NULL,
  "passwordHash" character varying(255) NOT NULL,
  "createdAt" timestamp NOT NULL,
  "updatedAt" timestamp NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "connections" table
CREATE TABLE "public"."connections" (
  "id" character varying(36) NOT NULL,
  "name" character varying(100) NOT NULL,
  "owner_id" character varying(36) NOT NULL,
  "createdAt" timestamp NOT NULL,
  "updatedAt" timestamp NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "owner_id" FOREIGN KEY ("owner_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
