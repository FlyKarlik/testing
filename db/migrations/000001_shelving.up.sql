CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO products (name) VALUES ('Ноутбук'),('Телевизор'),('Телефон'),('Системный блок'),('Часы'),('Микрофон');

CREATE TABLE IF NOT EXISTS shelves (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO shelves(name) VALUES ('Стеллаж А'),('Стеллаж Б'),('Стеллаж Ж');

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL NOT NULL,
    order_id INTEGER PRIMARY KEY
);

INSERT INTO orders (order_id) VALUES (10),(11),(14),(15);

CREATE TABLE IF NOT EXISTS products_orders (
    id SERIAL PRIMARY KEY,
    id_product INTEGER,
    id_order INTEGER,
    quantity INTEGER,
    FOREIGN KEY (id_product) REFERENCES products(id),
    FOREIGN KEY (id_order) REFERENCES orders(order_id)
);

INSERT INTO products_orders(id_product, id_order, quantity) VALUES (1, 10, 2),(2, 11, 3),(1, 14, 3),(3, 10, 1),(4, 14, 4),(5, 15, 1),(6, 10, 1);

CREATE TABLE IF NOT EXISTS products_shelves (
    id SERIAL PRIMARY KEY,
    id_product INTEGER,
    id_shelf INTEGER,
    is_main BOOLEAN,
    additional_shelves VARCHAR(255),
    FOREIGN KEY (id_product) REFERENCES products(id),
    FOREIGN KEY (id_shelf) REFERENCES shelves(id)
);

INSERT INTO products_shelves (id_product, id_shelf, is_main, additional_shelves) VALUES 
(1, 1, TRUE, ''),
(2, 1, TRUE, ''),
(3, 2, FALSE, 'З,В'),
(4, 3, TRUE, ''),
(5, 3, FALSE, 'А'),
(6, 3, TRUE, '');





