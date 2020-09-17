CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);

-- CREATE TABLE IF NOT EXISTS lists
-- (
--     id INTEGER PRIMARY KEY,
--     name varchar,
--
--
-- );