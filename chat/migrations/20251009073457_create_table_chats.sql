-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS chats (
    id              UUID NOT NULL DEFAULT gen_random_uuid(),         -- Уникальный идентификатор чата
    creator_id      UUID NOT NULL,                                   -- Идентификатор создателя чата
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Дата создания чата

    PRIMARY KEY(id)
);

COMMENT ON TABLE chats IS 'Таблица чатов';

COMMENT ON COLUMN chats.id IS 'Уникальный идентификатор чата';
COMMENT ON COLUMN chats.creator_id IS 'Идентификатор создателя чата';
COMMENT ON COLUMN chats.created_at IS 'Дата создания чата';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chats;
-- +goose StatementEnd
