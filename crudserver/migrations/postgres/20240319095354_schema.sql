-- +goose Up
-- +goose StatementBegin
create table if not exists projects (
    id serial primary key,
    name text not null,
    created_at timestamp with time zone not null default now()
);
create index projects_idx_project_id on projects(id);

insert into projects (name) values ('default') on conflict do nothing;

create table if not exists goods (
    id serial,
    project_id int not null,
        primary key (id, project_id),
    name text not null,
    description text,
    priority int not null,
    removed boolean not null default false,
    created_at timestamp with time zone not null default now()
);

create index goods_idx_good_id on goods(id);
create index goods_idx_project_id on goods(project_id);
create index goods_idx_name on goods(name);

create or replace function set_priority() returns trigger as $set_priority$
    declare max_priority int default 0;
    begin
        select from goods max(priority) into max_priority;
        new.priority = max_priority + 1;
    end;
$set_priority$ language plpgsql;

create trigger set_priority before insert on goods
for each row execute procedure set_priority();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists projects;
drop table if exists goods;
-- +goose StatementEnd
