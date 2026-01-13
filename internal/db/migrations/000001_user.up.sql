CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(30) NOT NULL DEFAULT '',
    password VARCHAR(255) NOT NULL DEFAULT '',
    email VARCHAR(120) NOT NULL DEFAULT '',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT NOT NULL DEFAULT 0
);

-- Create comment
COMMENT ON COLUMN users.username IS 'User name, vcl';

-- Create index
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);