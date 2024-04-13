package repository

import (
	"chat-app/pkg/models"

	"github.com/jmoiron/sqlx"
)

type MessageRepository interface {
	Create(message *models.Message) error
	GetMessages(senderID, recipientID uint) ([]models.Message, error)
}

type messageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *models.Message) error {
	const query = `
    INSERT INTO messages (sender_id, recipient_id, content, timestamp)
    VALUES ($1, $2, $3, $4)
    RETURNING id`
	err := r.db.QueryRow(query, message.SenderID, message.RecipientID, message.Content, message.Timestamp).Scan(&message.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *messageRepository) GetMessages(senderID, recipientID uint) ([]models.Message, error) {
	const query = `
    SELECT id, sender_id, recipient_id, content, timestamp
    FROM messages
    WHERE (sender_id = $1 AND recipient_id = $2) OR (sender_id = $2 AND recipient_id = $1)
    ORDER BY timestamp ASC`
	var messages []models.Message
	err := r.db.Select(&messages, query, senderID, recipientID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
