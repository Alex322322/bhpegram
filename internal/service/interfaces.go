package service

import (
	"context"

	"github.com/Alex322322/bhpegram/internal/domain"
)

type UserService interface {
	Create(ctx context.Context, user *domain.CreateUserRequest) (*domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id int64) error
}

type ChatService interface {
	//Create(ctx context.Context, chat *domain.Chat) (*domain.Chat, error)
	CreateWithMember(ctx context.Context, chat *domain.Chat, member *domain.ChatMember) (*domain.Chat, error)
	GetByID(ctx context.Context, id int64) (*domain.Chat, error)
	GetUserChats(ctx context.Context, userID int64) ([]*domain.Chat, error)
	AddMember(ctx context.Context, member *domain.ChatMember) error
	RemoveMember(ctx context.Context, chatID int64, userID int64) error
	Update(ctx context.Context, chat *domain.Chat) (*domain.Chat, error)
	Delete(ctx context.Context, id int64) error
}

type MessageService interface {
	Create(ctx context.Context, message *domain.CreateMessageRequest) (*domain.Message, error)
	GetByID(ctx context.Context, id int64) (*domain.Message, error)
	Update(ctx context.Context, message *domain.Message) (*domain.Message, error)
	Delete(ctx context.Context, id int64) error
	AddReaction(ctx context.Context, reaction *domain.MessageReaction) error
	UpdateReaction(ctx context.Context, reaction *domain.MessageReaction) error
	DeleteReaction(ctx context.Context, reaction *domain.MessageReaction) error
	GetReactions(ctx context.Context, id int64) ([]*domain.MessageReaction, error)
	GetRead(ctx context.Context, id int64) ([]*domain.MessageRead, error)
	// beforeID is used for cursor-based pagination, it returns messages with ID less than beforeID
	GetMessages(ctx context.Context, chatID int64, beforeID int64, limit int) ([]*domain.Message, error)
}