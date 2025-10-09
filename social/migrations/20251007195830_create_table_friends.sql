-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS friends (
    user_id    UUID NOT NULL,                                   -- Идентификатор пользователя
    friend_id  UUID NOT NULL,                                   -- Идентификатор друга пользователя
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Дата создания записи
    
    CHECK (user_id < friend_id) -- гарантирует уникальность записи и неравенство user_id != friend_id
);

COMMENT ON TABLE friends IS 'Таблица друзей пользователей';

COMMENT ON COLUMN friends.user_id IS 'Идентификатор пользователя';
COMMENT ON COLUMN friends.friend_id IS 'Идентификатор друга пользователя';
COMMENT ON COLUMN friends.created_at IS 'Дата создания записи';

-- Индексы
CREATE INDEX idx_user_friend ON friends(user_id, friend_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friends;
-- +goose StatementEnd
