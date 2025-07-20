-- Drop indexes first
DROP INDEX IF EXISTS idx_group_events_group_id;
DROP INDEX IF EXISTS idx_group_events_creator_id;
DROP INDEX IF EXISTS idx_group_events_event_date;

-- Drop the group_events table
DROP TABLE IF EXISTS group_events;
