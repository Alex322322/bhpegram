package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Alex322322/bhpegram/internal/domain"
	"github.com/Alex322322/bhpegram/internal/repository"
)

var _ MessageService = (*messageService)(nil)

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) *messageService {
	return &messageService{repo: repo}
}

func (s *messageService) Create(ctx context.Context, message *domain.CreateMessageRequest) (*domain.Message, error) {
	m, err := s.repo.Create(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}
	return m, nil
}

func (s *messageService) GetByID(ctx context.Context, id int64) (*domain.Message, error) {
	m, err := s.repo.GetByID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get message by ID: %w", err)
	}
	return m, nil
}

func (s *messageService) Update(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	m, err := s.repo.Update(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("failed to update message: %w", err)
	}
	return m, nil
}

func (s *messageService) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	return nil
}

func (s *messageService) AddReaction(ctx context.Context, reaction *domain.MessageReaction) error {
	err := s.repo.AddReaction(ctx, reaction)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to add reaction to message: %w", err)
	}
	return nil
}

func (s *messageService) UpdateReaction(ctx context.Context, reaction *domain.MessageReaction) error {
	err := s.repo.UpdateReaction(ctx, reaction)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to update reaction to message: %w", err)
	}
	return nil
}

func (s *messageService) DeleteReaction(ctx context.Context, reaction *domain.MessageReaction) error {
	err := s.repo.DeleteReaction(ctx, reaction)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to delete reaction to message: %w", err)
	}
	return nil
}

func (s *messageService) GetReactions(ctx context.Context, id int64) ([]*domain.MessageReaction, error) {
	reactions, err := s.repo.GetReactions(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get message reactions: %w", err)
	}
	return reactions, nil
}

func (s *messageService) GetRead(ctx context.Context, id int64) ([]*domain.MessageRead, error) {
	reads, err := s.repo.GetRead(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get message reads: %w", err)
	}
	return reads, nil
}

func (s *messageService) GetMessages(ctx context.Context, chatID int64, beforeID int64, limit int) ([]*domain.Message, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	messages, err := s.repo.GetMessages(ctx, chatID, beforeID, limit)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	return messages, nil
}
