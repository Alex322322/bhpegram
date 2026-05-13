package domain

import "time"

type ChatType string

const (
	Direct ChatType = "direct"
	Group  ChatType = "group"
)

type Chat struct {
	ID          int64
	Type        ChatType
	Name        string
	Description string
	PhotoURL    string
	CreatedBy   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChatMember struct {
	ChatID   int64
	UserID   int64
	Role     string
	JoinedAt time.Time
}
