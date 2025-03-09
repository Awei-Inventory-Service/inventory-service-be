CREATE TABLE sales (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_id UUID NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    -- product_id UUID,
    type VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    cost DECIMAL NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Foreign key to branch_id
    CONSTRAINT fk_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL
    -- CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products(uuid)
    -- ON UPDATE CASCADE
    -- ON DELETE SET NULL
);