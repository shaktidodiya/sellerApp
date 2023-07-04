USE SellerApp;

CREATE TABLE orders (
    id integer PRIMARY KEY AUTO_INCREMENT,
    user_id varchar(20),
    product_id varchar(20),
    price float,
    qty integer,
    created_at TIMESTAMP
);