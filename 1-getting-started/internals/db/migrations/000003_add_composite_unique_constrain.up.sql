-- Add a unique constraint on user_id + title
ALTER TABLE notes
ADD CONSTRAINT unique_user_note UNIQUE (user_id, title);