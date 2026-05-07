CREATE TYPE chat_type AS ENUM ('direct', 'group');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    phone_number VARCHAR(15) UNIQUE,
    email VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    type chat_type NOT NULL DEFAULT 'group',
    name VARCHAR(100),
    desciption VARCHAR(500),
    photo_url VARCHAR(500),
    created_by BIGINT REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE chat_members(
    chat_id BIGINT REFERENCES chats(id) ON DELETE CASCADE,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(20) NOT NULL DEFAULT 'member', -- 'admin', 'member', 'owner'
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (chat_id, user_id)
);

CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    author_id BIGINT REFERENCES users(id),
    content TEXT NOT NULL,
    message_type VARCHAR(20) NOT NULL DEFAULT 'text', -- 'text', 'image', 'file', 'system'.
    is_edited BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    edited_at TIMESTAMPTZ
);

CREATE TABLE message_reads (
    message_id BIGINT      NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
    user_id    BIGINT      NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    read_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (message_id, user_id)
);

CREATE TABLE reactions (
    id         BIGSERIAL    PRIMARY KEY,
    name       VARCHAR(50)  NOT NULL UNIQUE, -- "like", "heart", "custom_1"
    image_url  VARCHAR(500),                 -- NULL если это стандартный emoji
    emoji      VARCHAR(10),                  -- NULL если это кастомная картинка
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT reactions_type_check CHECK (
        (emoji IS NOT NULL AND image_url IS NULL) OR
        (emoji IS NULL AND image_url IS NOT NULL)
    )
);

CREATE TABLE message_reactions (
    message_id  BIGINT      NOT NULL REFERENCES messages(id)  ON DELETE CASCADE,
    user_id     BIGINT      NOT NULL REFERENCES users(id)     ON DELETE CASCADE,
    reaction_id BIGINT      NOT NULL REFERENCES reactions(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (message_id, user_id, reaction_id)
);