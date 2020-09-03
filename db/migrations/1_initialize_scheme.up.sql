CREATE TABLE IF NOT EXISTS users (
     id BIGSERIAL PRIMARY KEY,
     name varchar(20),
     email varchar(50) unique NOT NULL
)