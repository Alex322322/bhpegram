package domain

import "time"

type ChatType string

const (
	Direct ChatType = "direct"
	Group  ChatType = "group"
)

type Chat struct {
	ID          int64     `db:"id"`
	Type        ChatType  `db:"type"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	PhotoURL    string    `db:"photo_url"`
	CreatedBy   int64     `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ChatMember struct {
	ChatID   int64     `db:"chat_id"`
	UserID   int64     `db:"user_id"`
	Role     string    `db:"role"`
	JoinedAt time.Time `db:"joined_at"`
}
