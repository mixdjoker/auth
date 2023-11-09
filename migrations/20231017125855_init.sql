-- +goose Up
-- +goose StatementBegin
create table public.users (
    user_id bigserial,
    name text not null,
    email text not null,
    password char(64) not null,
    role_id int not null default 1,
    created_at timestamp not null default now(),
    updated_at timestamp,
    constraint users_pk primary key ("user_id"),
    constraint users_email_un unique ("email")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table public.users;
-- +goose StatementEnd
