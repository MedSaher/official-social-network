-- Drop indexes first
DROP INDEX IF EXISTS idx_group_invitations_group_id;
DROP INDEX IF EXISTS idx_group_invitations_invitee_id;
DROP INDEX IF EXISTS idx_group_invitations_status;

-- Drop the group_invitations table
DROP TABLE IF EXISTS group_invitations;
