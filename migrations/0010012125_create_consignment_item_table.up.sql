CREATE TABLE consignment_items (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    cost DECIMAL NOT NULL,
    price DECIMAL NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Foreign key to branch_id
    CONSTRAINT fk_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);