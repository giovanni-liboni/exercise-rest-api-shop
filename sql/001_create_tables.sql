# Language SQL
# Dialect SQL: mysql

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    role ENUM('admin', 'user'),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE INDEX (username)
);

CREATE TABLE IF NOT EXISTS items (
     id INT PRIMARY KEY AUTO_INCREMENT,
     producer VARCHAR(255) NOT NULL,
     description TEXT NOT NULL,
     price FLOAT NOT NULL,
     category VARCHAR(255) NOT NULL,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    payment_method ENUM('cash', 'card', 'paypal', 'stripe'),
    payment_id VARCHAR(255),
    total_price FLOAT NOT NULL,
    status ENUM('pending', 'paid', 'cancelled'),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS orders_items (
    order_id INT NOT NULL,
    item_id INT NOT NULL,
    price FLOAT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items(id)
);

ALTER TABLE users AUTO_INCREMENT=0;
ALTER TABLE items AUTO_INCREMENT=0;
ALTER TABLE orders AUTO_INCREMENT=0;