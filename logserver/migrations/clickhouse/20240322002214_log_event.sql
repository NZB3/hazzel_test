-- +goose Up
-- +goose StatementBegin
create table if not exists log (
    id Int,
    time DateTime,
    message String
)
    ENGINE = MergeTree
    ORDER BY (time, id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists log;
-- +goose StatementEnd
