-- +goose Up
-- +goose StatementBegin

-- Создаем партиции для outbox_events
-- Партиция для прошлых данных (до 2025)
CREATE TABLE IF NOT EXISTS outbox_events_before_2025 PARTITION OF outbox_events
    FOR VALUES FROM (MINVALUE) TO ('2025-01-01 00:00:00+00');

-- Партиция для текущего года (2025)
CREATE TABLE IF NOT EXISTS outbox_events_2025 PARTITION OF outbox_events
    FOR VALUES FROM ('2025-01-01 00:00:00+00') TO ('2026-01-01 00:00:00+00');

-- Партиция для следующего года (2026)
CREATE TABLE IF NOT EXISTS outbox_events_2026 PARTITION OF outbox_events
    FOR VALUES FROM ('2026-01-01 00:00:00+00') TO ('2027-01-01 00:00:00+00');

-- Партиция для будущих данных (после 2026)
CREATE TABLE IF NOT EXISTS outbox_events_after_2026 PARTITION OF outbox_events
    FOR VALUES FROM ('2027-01-01 00:00:00+00') TO (MAXVALUE);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS outbox_events_before_2025;
DROP TABLE IF EXISTS outbox_events_2025;
DROP TABLE IF EXISTS outbox_events_2026;
DROP TABLE IF EXISTS outbox_events_after_2026;

-- +goose StatementEnd

