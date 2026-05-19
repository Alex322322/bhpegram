package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/Alex322322/bhpegram/internal/domain"
	"github.com/jmoiron/sqlx"
)

var _ MessageRepository = (*PostgresMessageRepository)(nil)

type PostgresMessageRepository struct {
	db *sqlx.DB
}

func NewPostgresMessageRepository(db *sqlx.DB) *PostgresMessageRepository {
	return &PostgresMessageRepository{db: db}
}

func (r *PostgresMessageRepository) Create(ctx context.Context, req *domain.CreateMessageRequest) (*domain.Message, error) {
	query := `
	INSERT INTO messages (chat_id, author_id, content, message_type)
	VALUES ($1, $2, $3, $4)
	RETURNING *`

	var message domain.Message
	err := r.db.QueryRowxContext(ctx, query, req.ChatID, req.AuthorID, req.Content, req.MessageType).StructScan(&message)
	if err != nil {
		return nil, fmt.Errorf("failed to insert message data to database: %w", err)
	}
	return &message, nil
}

func (r *PostgresMessageRepository) GetByID(ctx context.Context, id int64) (*domain.Message, error) {
	query := `SELECT * FROM messages WHERE id = $1`

	var message domain.Message
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&message)
	if err != nil {
		return nil, fmt.Errorf("failed to get message by ID: %w", err)
	}
	return &message, nil
}

func (r *PostgresMessageRepository) Update(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	query := `
	UPDATE messages
	SET content = $2, message_type = $3, is_edited = $4, edited_at = NOW()
	WHERE id = $1
	RETURNING *`

	var updatedMessage domain.Message
	err := r.db.QueryRowxContext(ctx, query, message.ID, message.Content, message.MessageType, message.IsEdited).StructScan(&updatedMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to update message from database: %w", err)
	}
	return &updatedMessage, nil
}

func (r *PostgresMessageRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM messages WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete message from database: %w", err)
	}
	return nil
}

func (r *PostgresMessageRepository) AddReaction(ctx context.Context, reaction *domain.MessageReaction) error {
	query := `
	INSERT INTO message_reactions (message_id, user_id, reaction_id)
	VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, reaction.MessageID, reaction.UserID, reaction.ReactionID)
	if err != nil {
		return fmt.Errorf("failed to insert message reaction data to database: %w", err)
	}
	return nil
}

func (r *PostgresMessageRepository) UpdateReaction(ctx context.Context, reaction *domain.MessageReaction) error {
	query := `
	UPDATE message_reactions
	SET reaction_id = $3
	WHERE message_id = $1 AND user_id = $2`

	_, err := r.db.ExecContext(ctx, query, reaction.MessageID, reaction.UserID, reaction.ReactionID)
	if err != nil {
		return fmt.Errorf("failed to update message reaction from database: %w", err)
	}
	return nil
}

func (r *PostgresMessageRepository) DeleteReaction(ctx context.Context, reaction *domain.MessageReaction) error {
	query := `DELETE FROM message_reactions WHERE message_id = $1 AND user_id = $2`

	_, err := r.db.ExecContext(ctx, query, reaction.MessageID, reaction.UserID)
	if err != nil {
		return fmt.Errorf("failed to delete message reaction from database: %w", err)
	}
	return nil
}

func (r *PostgresMessageRepository) GetReactions(ctx context.Context, id int64) ([]*domain.MessageReaction, error) {
	query := `
	SELECT message_id, user_id, reaction_id, created_at FROM message_reactions
	WHERE message_reactions.message_id = $1`

	var reactions []*domain.MessageReaction
	err := r.db.SelectContext(ctx, &reactions, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get message reactions: %w", err)
	}
	return reactions, nil
}

func (r *PostgresMessageRepository) GetRead(ctx context.Context, id int64) ([]*domain.MessageRead, error) {
	query := `SELECT message_id, user_id, read_at FROM message_reads
	WHERE message_reads.message_id = $1`

	var reads []*domain.MessageRead
	err := r.db.SelectContext(ctx, &reads, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get message reads: %w", err)
	}
	return reads, nil
}

func (r *PostgresMessageRepository) GetMessages(ctx context.Context, chatID int64, beforeID int64, limit int) ([]*domain.Message, error) {
	query := `
	SELECT * FROM messages 
	WHERE chat_id = $1 AND id < $2
	ORDER BY id DESC
	LIMIT $3`

	var messages []*domain.Message
	if beforeID == 0 {
		beforeID = math.MaxInt64
	}
	err := r.db.SelectContext(ctx, &messages, query, chatID, beforeID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages from database: %w", err)
	}
	return messages, nil
}
