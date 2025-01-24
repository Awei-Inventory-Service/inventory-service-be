CREATE TABLE stock_transactions (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_origin_id UUID NOT NULL,
    item_id UUID NOT NULL,
    branch_destination_id UUID NOT NULL,
    issuer_id UUID NOT NULL,
    type VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    cost DECIMAL NOT NULL,
    reference VARCHAR(255) NOT NULL,
    remarks TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Foreign key to branch_origin_id
    CONSTRAINT fk_branch_origin_id FOREIGN KEY (branch_origin_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Foreign key to item_id
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES items(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Foreign key to branch_destination_id
    CONSTRAINT fk_branch_destination_id FOREIGN KEY (branch_destination_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Foreign key to issuer_id
    CONSTRAINT fk_issuer_id FOREIGN KEY (issuer_id) REFERENCES users(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);