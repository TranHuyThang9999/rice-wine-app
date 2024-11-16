CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255),
    phone_number VARCHAR(20),
    email VARCHAR(255),
    password VARCHAR(255),
    role INT DEFAULT 0,
    created_at INT,
    updated_at INT
);

CREATE TABLE IF NOT EXISTS file_stores (
    id BIGINT PRIMARY KEY,
    any_id BIGINT NOT NULL,
    path VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS type_rices (
    id BIGINT PRIMARY KEY,
    creator_id BIGINT,
    name VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP
);
