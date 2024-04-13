package models

import "time"

type Message struct {
	ID          uint      `db:"id" json:"id"`
	SenderID    uint      `db:"sender_id" json:"sender_id"`
	RecipientID uint      `db:"recipient_id" json:"recipient_id"`
	Content     string    `db:"content" json:"content"`
	Timestamp   time.Time `db:"timestamp" json:"timestamp"`
}
