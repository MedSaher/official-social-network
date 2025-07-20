-- Drop indexes first
DROP INDEX IF EXISTS idx_comments_post_id;
DROP INDEX IF EXISTS idx_comments_user_id;
DROP INDEX IF EXISTS idx_comments_created_at;

-- Drop the comments table
DROP TABLE IF EXISTS comments;
