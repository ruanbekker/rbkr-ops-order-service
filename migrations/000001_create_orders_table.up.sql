CREATE TABLE orders (
    id UUID PRIMARY KEY,
    product_id TEXT NOT NULL,
    quantity INT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
