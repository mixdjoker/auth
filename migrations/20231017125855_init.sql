-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users (
    id serial4 not NULL,
    name varchar(255),
    email varchar(255) UNIQUE,
    password char(64),
    role int not NULL,
    create_at timestamp not NULL,
    update_at timestamp,
    CONSTRAINT users_pk PRIMARY KEY (id),
    CONSTRAINT users_email_un UNIQUE ("email")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table public.users;
-- +goose StatementEnd
