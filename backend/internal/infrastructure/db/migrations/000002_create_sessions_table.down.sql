-- Drop indexes explicitly (good practice)
DROP INDEX IF EXISTS idx_sessions_user_id;
DROP INDEX IF EXISTS idx_sessions_expires_at;

-- Drop the sessions table
DROP TABLE IF EXISTS sessions;
