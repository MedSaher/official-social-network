-- Drop indexes first
DROP INDEX IF EXISTS idx_private_messages_sender_id;
DROP INDEX IF EXISTS idx_private_messages_receiver_id;
DROP INDEX IF EXISTS idx_private_messages_created_at;
DROP INDEX IF EXISTS idx_private_messages_is_read;

-- Drop the private_messages table
DROP TABLE IF EXISTS private_messages;
