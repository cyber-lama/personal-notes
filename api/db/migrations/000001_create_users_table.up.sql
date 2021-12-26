CREATE TABLE IF NOT EXISTS users(
    id bigserial primary key unique,
    username VARCHAR (50) unique not null,
    password VARCHAR (50) not null,
    email VARCHAR (300) unique not null
);