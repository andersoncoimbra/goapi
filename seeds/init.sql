use api;

CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, price)
    VALUES 
        ('Produto 1', 10.00),
        ('Produto 2', 20.00),
        ('Produto 3', 30.00),
        ('Produto 4', 40.00),
        ('Produto 5', 50.00),
        ('Produto 6', 60.00),
        ('Produto 7', 70.00),
        ('Produto 8', 80.00),
        ('Produto 9', 90.00),
        ('Produto 10', 100.00);

