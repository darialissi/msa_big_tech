package outbox

const tableOutboxEvents = "outbox_events"

const (
	outboxTableColumnID            = "id"
	outboxTableColumnAggType       = "aggregate_type"
	outboxTableColumnAggID         = "aggregate_id"
	outboxTableColumnEventType     = "event_type"
	outboxTableColumnPayload       = "payload"
	outboxTableColumnCreatedAt     = "created_at"
	outboxTableColumnPublishedAt   = "published_at"
	outboxTableColumnRetryCount    = "retry_count"
	outboxTableColumnNextAttemptAt = "next_attempt_at"
)

var tableOutboxEventsColumns = []string{
	outboxTableColumnID,
	outboxTableColumnAggType,
	outboxTableColumnAggID,
	outboxTableColumnEventType,
	outboxTableColumnPayload,
	outboxTableColumnCreatedAt,
	outboxTableColumnPublishedAt,
	outboxTableColumnRetryCount,
	outboxTableColumnNextAttemptAt,
}
