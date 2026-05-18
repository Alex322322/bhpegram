package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Alex322322/bhpegram/internal/domain"
	"github.com/Alex322322/bhpegram/internal/repository"
)

var _ ChatService = (*chatService)(nil)

type chatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) *chatService {
	return &chatService{repo: repo}
}

func (s *chatService) Create(ctx context.Context, chat *domain.Chat) (*domain.Chat, error) {
	c, err := s.repo.Create(ctx, chat)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	err = s.repo.AddMember(ctx, &domain.ChatMember{
		ChatID: c.ID,
		UserID: c.CreatedBy,
		Role:   "owner",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add creator as member: %w", err)
	}

	return c, nil
}

func (s *chatService) GetByID(ctx context.Context, id int64) (*domain.Chat, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get chat by id: %w", err)
	}
	return c, nil
}

func (s *chatService) GetUserChats(ctx context.Context, userID int64) ([]*domain.Chat, error) {
	chats, err := s.repo.GetUserChats(ctx, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get user chats: %w", err)
	}
	return chats, nil
}

func (s *chatService) AddMember(ctx context.Context, member *domain.ChatMember) error {
	err := s.repo.AddMember(ctx, member)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to add member to chat: %w", err)
	}
	return nil
}

func (s *chatService) RemoveMember(ctx context.Context, chatID int64, userID int64) error {
	err := s.repo.RemoveMember(ctx, chatID, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to remove member from chat: %w", err)
	}
	return nil
}

func (s *chatService) Update(ctx context.Context, chat *domain.Chat) (*domain.Chat, error) {
	c, err := s.repo.Update(ctx, chat)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to update chat: %w", err)
	}
	return c, nil
}

func (s *chatService) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to delete chat: %w", err)
	}
	return nil
}
