-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users (
    id serial,
    name text not null,
    email text not null,
    password char(64) not null,
    role int not null default 1,
    create_at timestamp not null default now(),
    update_at timestamp,
    CONSTRAINT users_pk PRIMARY KEY (id),
    CONSTRAINT users_email_un UNIQUE ("email")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table public.users;
-- +goose StatementEnd
