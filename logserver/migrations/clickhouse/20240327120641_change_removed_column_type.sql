-- +goose Up
-- +goose StatementBegin
alter table goods_events modify column removed UInt8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table goods_events modify column removed Bool;
-- +goose StatementEnd
