CREATE TABLE "users"
(
    "id"           varchar(36)  NOT NULL,
    "name"         varchar(100) NOT NULL,
    "passwordHash" varchar(255) NOT NULL,
    "createdAt"    timestamp    NOT NULL,
    "updatedAt"    timestamp    NOT NULL,

    PRIMARY KEY ("id")
);

CREATE TABLE "connections"
(
    "id"        varchar(36)  NOT NULL,
    "name"      varchar(100) NOT NULL,
    "owner_id"  varchar(36)  NOT NULL,
    "createdAt" timestamp    NOT NULL,
    "updatedAt" timestamp    NOT NULL,

    PRIMARY KEY ("id"),
    CONSTRAINT "owner_id" FOREIGN KEY ("owner_id") REFERENCES "users" ("id")
);