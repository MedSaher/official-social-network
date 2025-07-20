-- Drop indexes first
DROP INDEX IF EXISTS idx_posts_user_id;
DROP INDEX IF EXISTS idx_posts_group_id;
DROP INDEX IF EXISTS idx_posts_privacy;
DROP INDEX IF EXISTS idx_posts_created_at;

-- Drop the posts table
DROP TABLE IF EXISTS posts;