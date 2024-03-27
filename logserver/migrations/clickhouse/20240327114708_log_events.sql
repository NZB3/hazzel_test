-- +goose Up
-- +goose StatementBegin
create table if not exists log_events (
   time DateTime,
   level String,
   message String
)
    ENGINE = MergeTree
    ORDER BY (time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists log;
-- +goose StatementEnd
