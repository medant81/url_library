create table authors
(
    id         serial not null
        constraint authors_pk
            primary key,
    first_name varchar(300),
    last_name  varchar(300),
    biography  varchar,
    birthday   date
);

alter table authors
    owner to admin;

create table books
(
    id        serial not null
        constraint books_pkey
            primary key,
    name      varchar(300),
    author_id integer,
    year      integer,
    isbn      varchar(50)
);

alter table books
    owner to admin;

create table users
(
    id         serial not null
        constraint users_pk
            primary key,
    user_name  varchar(50),
    first_name varchar(300),
    last_name  varchar(300),
    email      varchar(300),
    password   varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

alter table users
    owner to admin;


create or replace function trigger_set_timestamp() returns trigger
    language plpgsql
as
$$
BEGIN
  --NEW.updated_at = clock_timestamp();
  NEW.email = 'test';
RETURN NEW;
END;
$$
;

alter function trigger_set_timestamp() owner to admin;


create trigger upd_users
    before update
    on users
    for each row
    execute procedure trigger_set_timestamp();

