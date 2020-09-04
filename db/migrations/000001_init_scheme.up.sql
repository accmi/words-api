CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name varchar(50),
    email varchar(255) NOT NULL
);
