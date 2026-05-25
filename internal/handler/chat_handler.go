package handler

import "github.com/Alex322322/bhpegram/internal/service"

type ChatHandler struct {
	service service.ChatService
}

func NewChatHandler(service service.ChatService) *ChatHandler {
	return &ChatHandler{service: service}
}