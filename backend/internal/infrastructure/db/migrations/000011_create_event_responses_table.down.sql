-- Drop indexes first
DROP INDEX IF EXISTS idx_event_responses_event_id;
DROP INDEX IF EXISTS idx_event_responses_user_id;
DROP INDEX IF EXISTS idx_event_responses_response;

-- Drop the event_responses table
DROP TABLE IF EXISTS event_responses;
