CREATE TABLE users
(
    id            varchar(36)  NOT NULL,
    name          varchar(100) NOT NULL,
    password_hash varchar(255) NOT NULL,
    created_at    timestamp    NOT NULL,
    updated_at    timestamp    NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE sessions
(
    id            varchar(36) NOT NULL,
    user_id       varchar(36) NOT NULL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    ip_address    varchar(36) NOT NULL,
    user_agent    varchar(36) NOT NULL,
    last_activity timestamp   NOT NULL,

    PRIMARY KEY (id)
);
CREATE INDEX sessions_userId_key ON sessions (user_id);
CREATE INDEX sessions_lastActivity_key ON sessions (last_activity);

CREATE TABLE connections
(
    id         varchar(36)  NOT NULL,
    name       varchar(100) NOT NULL,
    owner_id   varchar(36)  NOT NULL,
    created_at timestamp    NOT NULL,
    updated_at timestamp    NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT owner_id FOREIGN KEY (owner_id) REFERENCES users (id)
);