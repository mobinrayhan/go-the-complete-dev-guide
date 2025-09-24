CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    username   VARCHAR(50)                         NOT NULL,
    email      VARCHAR(255)                        NOT NULL,
    password   VARCHAR(255)                        NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULl,
    CONSTRAINT email
        UNIQUE (email),
    CONSTRAINT username
        UNIQUE (username)
);

CREATE INDEX IF NOT EXISTS username
    ON users (username, email);

DROP TRIGGER IF EXISTS update_modified_column ON notes;
CREATE TRIGGER update_modified_time BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();