CREATE TABLE IF NOT EXISTS bjora.users (
    id BIGINT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    gender TEXT NOT NULL,
    address TEXT NOT NULL,
    birthday TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT NOT NULL DEFAULT 'bjora-api'::TEXT,
    last_modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_modified_by TEXT NOT NULL DEFAULT 'bjora-api'::TEXT,
    is_active BOOLEAN DEFAULT true
);

CREATE UNIQUE INDEX users_username_unique ON bjora.users(username) WHERE (is_active IS TRUE);
CREATE UNIQUE INDEX users_email_unique ON bjora.users(email) WHERE (is_active IS TRUE);
CREATE INDEX IF NOT EXISTS users_id_is_active ON bjora.users USING btree(id, is_active);
CREATE INDEX IF NOT EXISTS users_email_is_active ON bjora.users USING btree(email, is_active);