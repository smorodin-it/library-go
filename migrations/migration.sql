create table profile
(
    id         varchar(36)  not null primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    patronymic varchar(255),
    phone      varchar(11)  not null,
    address    text         not null,
    user_id    varchar(36)  not null references "user" (id) on delete restrict,
    created_at timestamp(0) default now(),
    updated_at timestamp(0) default now()
);

create table "user"
(
    id            varchar(36)  not null primary key,
    email         varchar(100) not null unique,
    password_hash text         not null,
    profile_id    varchar(36)  not null references profile (id) on delete restrict,
    active        bool         default false,
    created_at    timestamp(0) default now(),
    updated_at    timestamp(0) default now()
);

create table library
(
    id         varchar(36)  not null primary key,
    name       varchar(255) not null,
    address    varchar(2000) default NULL::varchar,
    active     bool          default false,
    created_at timestamp(0)  default now(),
    updated_at timestamp(0)  default now()
);

create table users_in_libraries
(
    id         varchar(36) not null primary key,
    user_id    varchar(36) not null references "user" (id) on delete restrict,
    library_id varchar(36) not null references library (id) on delete restrict,
    created_at timestamp(0) default now(),
    updated_at timestamp(0) default now()

);

create table book
(
    id         varchar(36)  not null primary key,
    title      varchar(255) not null,
    author     varchar(255) default NULL::varchar,
    active     bool         default false,
    created_at timestamp(0) default now(),
    updated_at timestamp(0) default now()
);

create table books_in_libraries
(
    id           varchar(36) not null primary key,
    library_id   varchar(36) not null references library (id) on delete restrict,
    book_id      varchar(36) not null references book (id) on delete restrict,
    amount_total integer     not null default 0,
    amount_fact  integer     not null default 0,
    created_at   timestamp(0)         default now(),
    updated_at   timestamp(0)         default now(),

    unique (library_id, book_id)
);

create table books_at_user
(
    id         varchar(36) not null primary key,
    user_id    varchar(36) not null references "user" (id) on delete restrict,
    book_id    varchar(36) not null references book (id) on delete restrict,
    created_at timestamp(0) default now(),
    updated_at timestamp(0) default now(),

    unique (user_id, book_id)
)