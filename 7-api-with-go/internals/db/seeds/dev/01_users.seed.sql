-- Insert into users
INSERT INTO users (id, username, email, password, created_at, updated_at) VALUES
                    (1, 'user1', 'user1@example.com', 'hashed_password_1', '2025-09-02 02:21:41', '2024-03-20 19:49:28'),
                    (2, 'user2', 'user2@example.com', 'hashed_password_2', '2024-01-13 09:45:35', '2024-02-05 11:37:31'),
                    (3, 'user3', 'user3@example.com', 'hashed_password_3', '2025-04-03 19:06:54', '2025-03-03 09:15:08'),
                    (4, 'user4', 'user4@example.com', 'hashed_password_4', '2025-07-26 23:36:51', '2024-06-26 19:02:44'),
                    (5, 'user5', 'user5@example.com', 'hashed_password_5', '2024-07-04 20:42:33', '2023-09-29 11:47:12'),
                    (6, 'user6', 'user6@example.com', 'hashed_password_6', '2024-12-19 18:18:36', '2024-02-28 06:25:16'),
                    (7, 'user7', 'user7@example.com', 'hashed_password_7', '2025-04-12 09:57:05', '2024-09-24 22:28:56'),
                    (8, 'user8', 'user8@example.com', 'hashed_password_8', '2024-12-13 12:35:48', '2024-11-16 03:53:53'),
                    (9, 'user9', 'user9@example.com', 'hashed_password_9', '2025-08-25 13:07:59', '2025-04-30 02:17:44'),
                    (10, 'user10', 'user10@example.com', 'hashed_password_10', '2025-04-29 02:42:56', '2025-04-07 20:33:27'),
                    (11, 'user11', 'user11@example.com', 'hashed_password_11', '2024-05-27 17:51:22', '2024-02-25 05:16:22'),
                    (12, 'user12', 'user12@example.com', 'hashed_password_12', '2024-01-26 18:30:24', '2023-09-27 09:41:32'),
                    (13, 'user13', 'user13@example.com', 'hashed_password_13', '2025-05-30 16:57:54', '2025-07-20 19:32:05'),
                    (14, 'user14', 'user14@example.com', 'hashed_password_14', '2024-05-11 02:54:09', '2024-12-26 01:30:56'),
                    (15, 'user15', 'user15@example.com', 'hashed_password_15', '2023-11-21 07:14:24', '2025-03-19 12:19:57')
ON DUPLICATE KEY UPDATE
    username = VALUES(username),
    email = VALUES(email),
    password = VALUES(password),
    updated_at = VALUES(updated_at);