-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users (
    user_id bigserial,
    name text not null,
    email text not null,
    password char(64) not null,
    role_id int not null default 1,
    created_at timestamp not null default now(),
    updated_at timestamp,
    CONSTRAINT users_pk PRIMARY KEY ("user_id"),
    CONSTRAINT users_email_un UNIQUE ("email")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table public.users;
-- +goose StatementEnd
