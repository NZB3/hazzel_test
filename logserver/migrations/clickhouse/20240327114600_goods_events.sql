-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goods_events (
    id Int,
    project_id Int,
    name String,
    description String,
    priority Int,
    removed Bool,
    event_time DateTime
) ENGINE = MergeTree()
      ORDER BY (event_time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists goods_events;
-- +goose StatementEnd
