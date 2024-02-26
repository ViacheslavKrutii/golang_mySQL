CREATE DATABASE IF NOT EXISTS orders;

CREATE TABLE orders (
    Username VARCHAR(255),
    orderID int NOT NULL AUTO_INCREMENT,
    order_body JSON,
    PRIMARY KEY (orderID)
);