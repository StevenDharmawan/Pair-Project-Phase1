CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255),
    `address` TEXT,
    phone_number INT,
    balance DECIMAL(10,2) DEFAULT 0,
    `role` VARCHAR(50) NOT NULL DEFAULT 'Customer',
    UNIQUE(email)
);

CREATE TABLE sizes (
    size_id INT PRIMARY KEY AUTO_INCREMENT,
    size_name VARCHAR(10) NOT NULL
);

CREATE TABLE products (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(255) NOT NULL,
    `description` TEXT
);

CREATE TABLE products_detail (
    product_detail_id INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT NOT NULL,
    size_id INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock INT DEFAULT 0,
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    FOREIGN KEY (size_id) REFERENCES sizes(size_id)
);

CREATE TABLE wishlists (
    wishlist_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    product_detail_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (product_detail_id) REFERENCES products_detail(product_detail_id)
);

CREATE TABLE carts (
    cart_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    product_detail_id INT NOT NULL,
    quantity INT NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (product_detail_id) REFERENCES products_detail(product_detail_id)
);

CREATE TABLE order_history (
    order_history_id INT PRIMARY KEY AUTO_INCREMENT,
    cart_id INT NOT NULL,
    OrderDate DATE NOT NULL,
    FOREIGN KEY (cart_id) REFERENCES carts(cart_id)
);

-- Insert Table

INSERT INTO products (product_name, `description`)
VALUES 
('T-Shirt', 'Cotton T-Shirt'), 
('Graphic T-Shirt', 'T-shirts that have some sort of image or other graphic design on them'), 
('Denim Jacket', 'a jacket that made from denim');

INSERT INTO sizes (size_name)
VALUES ('S'), ('M'), ('L'), ('XL');

INSERT INTO products_detail (product_id, size_id, price, stock)
VALUES 
(1, 1, 200000, 50),
(1, 2, 205000, 40),
(1, 3, 210000, 40),
(1, 4, 215000, 30),
(2, 1, 150000, 50),
(2, 2, 155000, 40),
(2, 3, 160000, 40),
(2, 4, 165000, 20),
(3, 1, 300000, 30),
(3, 2, 305000, 20),
(3, 3, 310000, 20),
(3, 4, 315000, 15);
