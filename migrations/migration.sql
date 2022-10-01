create table library
(
    id         varchar(36)  not null primary key,
    name       varchar(255) not null,
    address    varchar(2000) default NULL::varchar,
    active     bool          default false,
    created_at timestamp(0)  default now(),
    updated_at timestamp(0)  default now()
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
)