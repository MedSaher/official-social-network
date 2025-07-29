package models

type GroupEvent struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	CreatorID   int    `json:"creator_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
	CreatedAt   string `json:"created_at"`
}
