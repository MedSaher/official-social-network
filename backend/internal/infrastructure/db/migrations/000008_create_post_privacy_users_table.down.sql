-- Drop indexes first
DROP INDEX IF EXISTS idx_post_privacy_users_post_id;
DROP INDEX IF EXISTS idx_post_privacy_users_user_id;

-- Drop the post_privacy_users table
DROP TABLE IF EXISTS post_privacy_users;
