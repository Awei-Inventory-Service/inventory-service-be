CREATE TABLE items (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(255) NOT NULL,
    portion_size DECIMAL(10,2) NOT NULL DEFAULT 1.0,
    supplier_id UUID,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_supplier_id FOREIGN KEY(supplier_id) REFERENCES suppliers(uuid)
);