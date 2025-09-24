CREATE TABLE IF NOT EXISTS notes
(
    id SERIAL PRIMARY KEY,
    user_id    INT                                 NOT NULL,
    title      VARCHAR(255)                        NOT NULL,
    body       TEXT                                NOT NULL,
    tags       JSONB                                NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_notes_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS user_id
    ON notes (user_id);

DROP TRIGGER IF EXISTS update_modified_column ON notes;
CREATE TRIGGER update_modified_column
BEFORE UPDATE ON notes
FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();
