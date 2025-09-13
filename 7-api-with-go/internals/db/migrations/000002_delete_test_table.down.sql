-- auto-generated definition
CREATE TABLE test_table
(
    id          INT AUTO_INCREMENT
        PRIMARY KEY,
    name        VARCHAR(100)                        NOT NULL,
    description TEXT                                NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
);

