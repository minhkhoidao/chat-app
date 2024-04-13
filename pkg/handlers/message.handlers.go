package handlers

import (
	"chat-app/pkg/models"
	"chat-app/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatService services.ChatService
}

func NewChatHandler(chatService services.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (h *ChatHandler) SendMessage(c *gin.Context) {
	senderID, _ := strconv.ParseUint(c.Param("senderID"), 10, 32)
	recipientID, _ := strconv.ParseUint(c.Param("recipientID"), 10, 32)
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	err := h.chatService.SendMessage(uint(senderID), uint(recipientID), message.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func (h *ChatHandler) GetConversation(c *gin.Context) {
	senderID, _ := strconv.ParseUint(c.Param("senderID"), 10, 32)
	recipientID, _ := strconv.ParseUint(c.Param("recipientID"), 10, 32)

	messages, err := h.chatService.GetConversation(uint(senderID), uint(recipientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
