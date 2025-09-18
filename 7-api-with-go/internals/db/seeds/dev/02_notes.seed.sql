INSERT INTO notes (user_id, title, body, tags)
VALUES
-- User 1 notes
(1, 'Golang Basics', 'Learning about structs, slices, and interfaces in Go.', '["golang", "learning", "backend"]'),
(1, 'Daily Tasks', '1. Standup meeting\n2. Review PRs\n3. Work on API endpoints', '["tasks", "work", "api"]'),

-- User 2 notes
(2, 'Shopping List', 'Milk, Bread, Eggs, Coffee, Chicken', '["shopping", "personal"]'),
(2, 'Travel Plans', 'Planning a trip to Cox’s Bazar next month. Book hotels early.', '["travel", "plans", "personal"]'),

-- User 3 notes
(3, 'Project Ideas', '1. Task manager app\n2. AI chatbot for study help\n3. Personal finance tracker', '["projects", "ideas", "apps"]'),
(3, 'Book Highlights', 'Notes from "Clean Code": meaningful names, small functions, single responsibility principle.', '["books", "clean-code", "learning"]')
ON DUPLICATE KEY UPDATE 
id = VALUES(id)
