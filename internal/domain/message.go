package domain

import "time"

type Message struct {
	ID          int64
	ChatID      int64
	AuthorID    int64
	Content     string
	MessageType string
	IsEdited    bool
	CreatedAt   time.Time
	EditedAt    *time.Time // can be nil
}

type MessageRead struct {
	MessageID int64
	UserID    int64
	ReadAt    time.Time
}

type Reaction struct {
	ID        int64
	Name      string
	ImageURL  *string // can be nil
	Emoji     *string // can be nil
	CreatedAt time.Time
}

type MessageReaction struct {
	MessageID  int64
	UserID     int64
	ReactionID int64
	CreatedAt  time.Time
}
