CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null unique,
    encrypted_password varchar not null
);