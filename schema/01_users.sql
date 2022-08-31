CREATE TABLE IF NOT EXISTS users
(
    user_id          SERIAL PRIMARY KEY, 
    name             VARCHAR(255) NOT NULL,
    email            VARCHAR(255) NOT NULL,
    password         TEXT NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_pk ON users (user_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email_unique ON users (email);