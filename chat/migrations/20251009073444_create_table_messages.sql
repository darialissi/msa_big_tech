-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages (
    id           UUID NOT NULL DEFAULT gen_random_uuid(),         -- Уникальный идентификатор сообщения
    chat_id      UUID NOT NULL,                                   -- Идентификатор чата
    sender_id    UUID NOT NULL,                                   -- Идентификатор отправителя
    text         TEXT NOT NULL,                                   -- Текст сообщения
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Дата создания записи

    PRIMARY KEY(id)
);

COMMENT ON TABLE messages IS 'Таблица сообщений';

COMMENT ON COLUMN messages.id IS 'Уникальный идентификатор сообщения';
COMMENT ON COLUMN messages.chat_id IS 'Идентификатор чата';
COMMENT ON COLUMN messages.sender_id IS 'Идентификатор отправителя';
COMMENT ON COLUMN messages.text IS 'Текст сообщения';
COMMENT ON COLUMN messages.created_at IS 'Дата создания записи';

-- Индексы
CREATE INDEX idx_message ON messages(chat_id, created_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
