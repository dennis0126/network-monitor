-- Create "sessions" table
CREATE TABLE "public"."sessions"
(
    "id"           character varying(36) NOT NULL,
    "userId"       character varying(36) NOT NULL,
    "ipAddress"    character varying(36) NOT NULL,
    "userAgent"    character varying(36) NOT NULL,
    "lastActivity" timestamp             NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "sessions_userId_fkey" FOREIGN KEY ("userId") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "sessions_lastactivity_key" to table: "sessions"
CREATE INDEX "sessions_lastactivity_key" ON "public"."sessions" ("lastActivity");
-- Create index "sessions_userid_key" to table: "sessions"
CREATE INDEX "sessions_userid_key" ON "public"."sessions" ("userId");
