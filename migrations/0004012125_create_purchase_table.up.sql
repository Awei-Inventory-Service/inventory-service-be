CREATE TABLE purchases (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    supplier_id UUID NOT NULL,
    branch_id UUID NOT NULL,
    item_id UUID NOT NULL,
    quantity INT NOT NULL,
    purchase_cost DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Add foreign key to supplier
    CONSTRAINT fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES suppliers(uuid),
    -- Add foreign key to branch
    CONSTRAINT fk_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid),
    -- Add foreign key to item
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES items(uuid)
);