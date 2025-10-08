-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS friend_requests (
    id              UUID NOT NULL DEFAULT gen_random_uuid(),         -- Уникальный идентификатор запроса
    from_user_id    UUID NOT NULL,                                   -- Уникальный идентификатор отправителя запроса
    to_user_id      UUID NOT NULL,                                   -- Уникальный идентификатор получателя запроса
    status          INTEGER NOT NULL,                                -- Статус запроса
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Дата создания записи
    PRIMARY KEY(id),
    CHECK (from_user_id != to_user_id)
);

COMMENT ON TABLE friend_requests IS 'Таблица запросов В Друзья';

COMMENT ON COLUMN friend_requests.id IS 'Уникальный идентификатор запроса';
COMMENT ON COLUMN friend_requests.from_user_id IS 'Уникальный идентификатор отправителя запроса';
COMMENT ON COLUMN friend_requests.to_user_id IS 'Уникальный идентификатор получателя запроса';
COMMENT ON COLUMN friend_requests.status IS 'Статус запроса';
COMMENT ON COLUMN friend_requests.created_at IS 'Дата создания записи';

-- Индексы
CREATE INDEX idx_friend_requests_users ON friend_requests(from_user_id, to_user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friend_requests;
-- +goose StatementEnd
