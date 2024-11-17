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

CREATE TABLE  IF NOT EXISTS rices (
    id BIGINT PRIMARY KEY,                -- ID tự tăng, khóa chính
    creator_id BIGINT NOT NULL,           -- ID người tạo
    type_rice_id BIGINT NOT NULL,         -- ID loại gạo
    name VARCHAR(255) NOT NULL,           -- Tên gạo
    quantity INT NOT NULL DEFAULT 0,      -- Số lượng (mặc định là 0)
    price NUMERIC(10, 2) NOT NULL,        -- Giá (số thập phân với tối đa 2 chữ số sau dấu phẩy)
    origin VARCHAR(255),                  -- Xuất xứ gạo
    harvest_season INT,                   -- Mùa thu hoạch
    created_at BIGINT,   -- Ngày tạo (mặc định là thời gian hiện tại)
    updated_at BIGINT
);

CREATE TABLE  IF NOT EXISTS orders
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGINT         NOT NULL,
    status       VARCHAR(50)    NOT NULL,
    total_amount NUMERIC(10, 2) NOT NULL,
    created_at   BIGINT         NOT NULL, -- Unix timestamp
    updated_at   BIGINT         NOT NULL  -- Unix timestamp
);

CREATE TABLE  IF NOT EXISTS  order_items
(
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT         NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
    product_id BIGINT         NOT NULL,
    quantity   INT            NOT NULL,
    price      NUMERIC(10, 2) NOT NULL,
    total      NUMERIC(10, 2) NOT NULL,
    created_at BIGINT         NOT NULL, -- Unix timestamp
    updated_at BIGINT         NOT NULL  -- Unix timestamp
);
