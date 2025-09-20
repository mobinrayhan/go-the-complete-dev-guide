
-- Remove the unique constraint
ALTER TABLE notes
DROP CONSTRAINT IF EXISTS unique_user_note;
