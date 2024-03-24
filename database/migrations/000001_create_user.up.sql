CREATE TABLE IF NOT EXISTS tbl_user(
    uid SERIAL PRIMARY KEY,
    username VARCHAR(30) NOT NULL UNIQUE ,
    password VARCHAR(200) NOT NULL,
    is_admin BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_user_username ON tbl_user(username);
