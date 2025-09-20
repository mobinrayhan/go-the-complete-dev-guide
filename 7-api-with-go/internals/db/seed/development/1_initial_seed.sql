INSERT INTO users (username, email, password)
VALUES
  ('alice', 'alice@example.com', 'hashedpassword1'),
  ('bob', 'bob@example.com', 'hashedpassword2'),
  ('charlie', 'charlie@example.com', 'hashedpassword3'),
  ('david', 'david@example.com', 'hashedpassword4'),
  ('eve', 'eve@example.com', 'hashedpassword5'),
  ('frank', 'frank@example.com', 'hashedpassword6'),
  ('grace', 'grace@example.com', 'hashedpassword7'),
  ('heidi', 'heidi@example.com', 'hashedpassword8'),
  ('ivan', 'ivan@example.com', 'hashedpassword9'),
  ('judy', 'judy@example.com', 'hashedpassword10'),
  ('kate', 'kate@example.com', 'hashedpassword11'),
  ('leo', 'leo@example.com', 'hashedpassword12'),
  ('mallory', 'mallory@example.com', 'hashedpassword13'),
  ('nancy', 'nancy@example.com', 'hashedpassword14'),
  ('oliver', 'oliver@example.com', 'hashedpassword15')
ON CONFLICT (email) DO NOTHING;

INSERT INTO notes (user_id, title, body, tags)
VALUES
  ((SELECT id FROM users WHERE username='alice'), 'Shopping', 'Buy groceries and milk', '{"category":"personal"}'),
  ((SELECT id FROM users WHERE username='alice'), 'Workout', 'Go to gym at 6pm', '{"category":"health"}'),
  ((SELECT id FROM users WHERE username='bob'), 'Project Plan', 'Finish the project plan by Monday', '{"priority":"high"}'),
  ((SELECT id FROM users WHERE username='bob'), 'Meeting Notes', 'Discuss roadmap', '{"priority":"medium"}'),
  ((SELECT id FROM users WHERE username='charlie'), 'Book List', 'Read 3 books this month', '{"category":"personal"}'),
  ((SELECT id FROM users WHERE username='charlie'), 'Work Tasks', 'Complete report and submit', '{"priority":"high"}'),
  ((SELECT id FROM users WHERE username='david'), 'Holiday', 'Plan holiday trip', '{"category":"leisure"}'),
  ((SELECT id FROM users WHERE username='eve'), 'Shopping', 'Buy birthday gift', '{"category":"personal"}'),
  ((SELECT id FROM users WHERE username='frank'), 'Gym Routine', 'Leg day workout', '{"category":"health"}'),
  ((SELECT id FROM users WHERE username='grace'), 'Client Meeting', 'Prepare presentation slides', '{"priority":"high"}'),
  ((SELECT id FROM users WHERE username='heidi'), 'Grocery', 'Weekly grocery shopping', '{"category":"personal"}'),
  ((SELECT id FROM users WHERE username='ivan'), 'Study', 'Finish Go tutorial', '{"category":"education"}'),
  ((SELECT id FROM users WHERE username='judy'), 'Finance', 'Update budget sheet', '{"category":"finance"}'),
  ((SELECT id FROM users WHERE username='kate'), 'Doctor Appointment', 'Checkup at 4pm', '{"category":"health"}'),
  ((SELECT id FROM users WHERE username='leo'), 'Car Service', 'Oil change and tire check', '{"category":"personal"}'),
  ((SELECT id FROM users WHERE username='mallory'), 'Team Meeting', 'Discuss sprint goals', '{"priority":"medium"}'),
  ((SELECT id FROM users WHERE username='nancy'), 'Volunteer', 'Community service on Saturday', '{"category":"social"}'),
  ((SELECT id FROM users WHERE username='oliver'), 'Learn SQL', 'Practice advanced queries', '{"category":"education"}'),
  ((SELECT id FROM users WHERE username='alice'), 'Cooking', 'Try new pasta recipe', '{"category":"personal"}'),
  ((SELECT id FROM users WHERE username='bob'), 'Movie Night', 'Watch a movie with friends', '{"category":"leisure"}')
ON CONFLICT (user_id, title) DO NOTHING;
