package handler

import "github.com/Alex322322/bhpegram/internal/service"

type MessageHandler struct {
	service service.MessageService
}

func NewMessageHandler(service service.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}