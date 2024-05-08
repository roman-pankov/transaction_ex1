CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    balance INT NOT NULL DEFAULT 100,
    version_update INT NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL
);

INSERT INTO users (username, email) VALUES
('user1', 'user1@example.com');

INSERT INTO products (name, price) VALUES
('product1', 10);