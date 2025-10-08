-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS friends (
    user_id    UUID NOT NULL,                                   -- Уникальный идентификатор пользователя
    friend_id  UUID NOT NULL,                                   -- Уникальный идентификатор друга пользователя
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Дата создания записи
    UNIQUE(user_id, friend_id),
    CHECK (user_id != friend_id)
);

COMMENT ON TABLE friends IS 'Таблица друзей пользователей';

COMMENT ON COLUMN friends.user_id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN friends.friend_id IS 'Уникальный идентификатор друга пользователя';
COMMENT ON COLUMN friends.created_at IS 'Дата создания записи';

-- Индексы
CREATE INDEX idx_friends_user ON friends(user_id);
CREATE INDEX idx_friends_friend ON friends(friend_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friends;
-- +goose StatementEnd
