CREATE TABLE IF NOT EXISTS users(
    id bigserial primary key unique,
    username VARCHAR (50) unique,
    password VARCHAR (50) not null,
    email VARCHAR (300) unique not null,
    created_at      timestamp not null,
    updated_at      timestamp not null
);