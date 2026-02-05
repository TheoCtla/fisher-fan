-- Migration: Add authentication fields to users table
-- Date: 2026-02-05

-- Add password column (required for authentication)
ALTER TABLE users ADD COLUMN IF NOT EXISTS password VARCHAR(255) NOT NULL DEFAULT '';

-- Add refresh_token column (for JWT refresh tokens)
ALTER TABLE users ADD COLUMN IF NOT EXISTS refresh_token TEXT;

-- Add timestamps for tracking user creation and updates
ALTER TABLE users ADD COLUMN IF NOT EXISTS created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE users ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Create index on email for faster lookups during login
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Note: Existing users will have empty passwords and will need to set them
-- You may want to run a separate script to handle existing users
