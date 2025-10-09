-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id           UUID NOT NULL,                                   -- Уникальный идентификатор пользователя
    nickname     VARCHAR(50) NOT NULL,                            -- Уникальный никнейм пользователя
    avatar_url   VARCHAR(300),                                    -- Ссылка на аватарку
    bio          VARCHAR(500),                                    -- Био
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Дата создания профиля

    PRIMARY KEY(id),

    UNIQUE(nickname)
);

COMMENT ON TABLE users IS 'Таблица пользователей';

COMMENT ON COLUMN users.id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN users.nickname IS 'Уникальный никнейм пользователя';
COMMENT ON COLUMN users.avatar_url IS 'Ссылка на аватарку';
COMMENT ON COLUMN users.bio IS 'Био';
COMMENT ON COLUMN users.created_at IS 'Дата создания записи';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
