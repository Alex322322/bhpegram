package domain

import "time"

type Message struct {
	ID          int64      `db:"id"`
	ChatID      int64      `db:"chat_id"`
	AuthorID    int64      `db:"author_id"`
	Content     string     `db:"content"`
	MessageType string     `db:"message_type"`
	IsEdited    bool       `db:"is_edited"`
	CreatedAt   time.Time  `db:"created_at"`
	EditedAt    *time.Time `db:"edited_at"` // can be nil
}

type MessageRead struct {
	MessageID int64     `db:"message_id"`
	UserID    int64     `db:"user_id"`
	ReadAt    time.Time `db:"read_at"`
}

type Reaction struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	ImageURL  *string   `db:"image_url"` // can be nil
	Emoji     *string   `db:"emoji"`     // can be nil
	CreatedAt time.Time `db:"created_at"`
}

type MessageReaction struct {
	MessageID  int64     `db:"message_id"`
	UserID     int64     `db:"user_id"`
	ReactionID int64     `db:"reaction_id"`
	CreatedAt  time.Time `db:"created_at"`
}

type CreateMessageRequest struct {
	ChatID      int64  `json:"chat_id"`
	AuthorID    int64  `json:"author_id"`
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
}
