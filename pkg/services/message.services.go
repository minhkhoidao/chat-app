package services

import (
	"chat-app/pkg/models"
	"chat-app/pkg/repository"
	"time"
)

type ChatService interface {
	SendMessage(senderID, recipientID uint, content string) error
	GetConversation(senderID, recipientID uint) ([]models.Message, error)
}

type chatService struct {
	messageRepo repository.MessageRepository
}

func NewChatService(messageRepo repository.MessageRepository) ChatService {
	return &chatService{messageRepo: messageRepo}
}

func (s *chatService) SendMessage(senderID, recipientID uint, content string) error {
	message := models.Message{
		SenderID:    senderID,
		RecipientID: recipientID,
		Content:     content,
		Timestamp:   time.Now(),
	}
	return s.messageRepo.Create(&message)
}

func (s *chatService) GetConversation(senderID, recipientID uint) ([]models.Message, error) {
	return s.messageRepo.GetMessages(senderID, recipientID)
}
