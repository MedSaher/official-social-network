-- Drop indexes first
DROP INDEX IF EXISTS idx_group_messages_group_id;
DROP INDEX IF EXISTS idx_group_messages_sender_id;
DROP INDEX IF EXISTS idx_group_messages_created_at;

-- Drop the group_messages table
DROP TABLE IF EXISTS group_messages;
