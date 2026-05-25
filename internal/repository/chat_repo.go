package repository

import (
	"context"
	"fmt"

	"github.com/Alex322322/bhpegram/internal/domain"
	"github.com/jmoiron/sqlx"
)

var _ ChatRepository = (*PostgresChatRepository)(nil)

type PostgresChatRepository struct {
	db *sqlx.DB
}

func NewPostgresChatRepository(db *sqlx.DB) *PostgresChatRepository {
	return &PostgresChatRepository{db: db}
}

/*
func (r *PostgresChatRepository) Create(ctx context.Context, req *domain.Chat) (*domain.Chat, error) {
	query := `
	INSERT INTO chats (type, name, description, photo_url, created_by)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *`

	var chat domain.Chat
	err := r.db.QueryRowxContext(ctx, query, req.Type, req.Name, req.Description, req.PhotoURL, req.CreatedBy).StructScan(&chat)
	if err != nil {
		return nil, fmt.Errorf("failed to insert chat data to database: %w", err)
	}
	return &chat, nil
}
*/

func (r *PostgresChatRepository) CreateWithMember(ctx context.Context, chat *domain.Chat) (*domain.Chat, error) {
    tx, err := r.db.BeginTxx(ctx, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback()

    query := `
    INSERT INTO chats (type, name, description, photo_url, created_by)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *`

    var createdChat domain.Chat
    err = tx.QueryRowxContext(ctx, query, chat.Type, chat.Name, chat.Description, chat.PhotoURL, chat.CreatedBy).StructScan(&createdChat)
    if err != nil {
        return nil, fmt.Errorf("failed to create chat: %w", err)
    }

    memberQuery := `
    INSERT INTO chat_members (chat_id, user_id, role)
    VALUES ($1, $2, $3)`

    _, err = tx.ExecContext(ctx, memberQuery, createdChat.ID, createdChat.CreatedBy, "owner")
    if err != nil {
        return nil, fmt.Errorf("failed to add creator as member: %w", err)
    }

    if err = tx.Commit(); err != nil {
        return nil, fmt.Errorf("failed to commit transaction: %w", err)
    }

    return &createdChat, nil
}

func (r *PostgresChatRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM chats WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete chat from database: %w", err)
	}
	return nil
}

func (r *PostgresChatRepository) Update(ctx context.Context, chat *domain.Chat) (*domain.Chat, error) {
	query := `
	UPDATE chats
	SET name = $2, description = $3, photo_url = $4, updated_at = NOW()
	WHERE id = $1
	RETURNING *`

	var updatedChat domain.Chat
	err := r.db.QueryRowxContext(ctx, query, chat.ID, chat.Name, chat.Description, chat.PhotoURL).StructScan(&updatedChat)
	if err != nil {
		return nil, fmt.Errorf("failed to update chat from database: %w", err)
	}
	return &updatedChat, nil
}

func (r *PostgresChatRepository) GetByID(ctx context.Context, id int64) (*domain.Chat, error) {
	query := `SELECT * FROM chats WHERE id = $1`

	var chat domain.Chat
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&chat)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat by ID: %w", err)
	}
	return &chat, nil
}

func (r *PostgresChatRepository) GetUserChats(ctx context.Context, userID int64) ([]*domain.Chat, error) {
	query := `
	SELECT id, type, name, description, photo_url, created_by, created_at, updated_at FROM chats 
	INNER JOIN chat_members ON chats.id = chat_members.chat_id
	WHERE chat_members.user_id = $1
	`

	var chat []*domain.Chat
	err := r.db.SelectContext(ctx, &chat, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user chats: %w", err)
	}
	return chat, nil
}

func (r *PostgresChatRepository) AddMember(ctx context.Context, member *domain.ChatMember) error {
	query := `
	INSERT INTO chat_members (chat_id, user_id, role, joined_at)
	VALUES ($1, $2, $3, NOW())`
	
	_, err := r.db.ExecContext(ctx, query, member.ChatID, member.UserID, member.Role)
	if err != nil {
		return fmt.Errorf("failed to insert chat member data to database: %w", err)
	}
	return nil
}

func (r *PostgresChatRepository) RemoveMember(ctx context.Context, chatID int64, userID int64) error {
	query := `DELETE FROM chat_members WHERE chat_id = $1 AND user_id = $2`
	
	_, err := r.db.ExecContext(ctx, query, chatID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete chat member from database: %w", err)
	}
	return nil
}
