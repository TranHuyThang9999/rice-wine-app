CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255),
    phone_number VARCHAR(20),
    password VARCHAR(255),
    role INT DEFAULT 0,
    created_at INT,
    updated_at INT
);

CREATE TABLE file_stores (
    id BIGINT PRIMARY KEY,
    any_id BIGINT NOT NULL,       
    path VARCHAR(255) NOT NULL   
);
