-- Drop indexes first (good migration practice)
DROP INDEX IF EXISTS idx_follows_follower_id;
DROP INDEX IF EXISTS idx_follows_following_id;
DROP INDEX IF EXISTS idx_follows_status;

-- Drop the follows table
DROP TABLE IF EXISTS follows;
