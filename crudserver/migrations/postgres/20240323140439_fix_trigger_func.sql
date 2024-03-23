-- +goose Up
-- +goose StatementBegin
create or replace function set_priority() returns trigger as $set_priority$
declare
    max_priority int;
begin
    select max(priority) into max_priority from goods;
    if max_priority is null then
        new.priority := 1; -- If no records exist, set priority to 1
    else
        new.priority := max_priority + 1;
    end if;
    return new;
end;
$set_priority$ language plpgsql;

create or replace trigger set_priority before insert on goods
    for each row execute function set_priority();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop trigger set_priority on goods;
-- +goose StatementEnd
