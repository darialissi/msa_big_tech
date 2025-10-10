-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_users (
    id              UUID NOT NULL DEFAULT gen_random_uuid(),         -- Уникальный идентификатор пользователя
    email           VARCHAR(100) NOT NULL,                           -- Уникальный Email пользователя
    password_hash   VARCHAR(500) NOT NULL,                           -- Хэшированный пароль пользователя
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(), -- Дата создания записи

    PRIMARY KEY(id),

    UNIQUE(email)
);

COMMENT ON TABLE auth_users IS 'Таблица зарегистрированных пользователей';

COMMENT ON COLUMN auth_users.id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN auth_users.email IS 'Уникальный Email пользователя';
COMMENT ON COLUMN auth_users.password_hash IS 'Хэшированный пароль пользователя';
COMMENT ON COLUMN auth_users.created_at IS 'Дата создания записи';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth_users;
-- +goose StatementEnd
