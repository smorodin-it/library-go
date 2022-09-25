create table library
(
    id         varchar(36)  not null primary key,
    name       varchar(255) not null,
    address    varchar(2000) default NULL::varchar,
    active     bool          default false,
    created_at timestamp(0)  default now(),
    updated_at timestamp(0)  default now()
)