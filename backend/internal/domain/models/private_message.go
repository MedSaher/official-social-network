package models

import "time"

type PrivateMessage struct {
    ID         int        `json:"id"`
    SenderID   int        `json:"sender_id"`
    ReceiverID int        `json:"receiver_id"`
    Content    string     `json:"content"`
    IsRead     bool       `json:"is_read"`
    CreatedAt  time.Time  `json:"created_at"`
}
