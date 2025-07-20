-- Drop indexes first
DROP INDEX IF EXISTS idx_group_members_group_id;
DROP INDEX IF EXISTS idx_group_members_user_id;
DROP INDEX IF EXISTS idx_group_members_status;

-- Drop the group_members table
DROP TABLE IF EXISTS group_members;
