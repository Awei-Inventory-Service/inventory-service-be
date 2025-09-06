CREATE TABLE branch_items (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_id UUID NOT NULL,
    item_id UUID NOT NULL,
    current_stock DECIMAL(10,2) NOT NULL,
    price DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Add foreign key to branch
    CONSTRAINT fk_branch_items_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
    
    -- Add foreign key to item
    CONSTRAINT fk_branch_items_item_id FOREIGN KEY (item_id) REFERENCES items(uuid)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
    
    -- Unique constraint for branch-item combination
    UNIQUE(branch_id, item_id)
);

-- Create indexes for better performance
CREATE INDEX idx_branch_items_branch_id ON branch_items(branch_id);
CREATE INDEX idx_branch_items_item_id ON branch_items(item_id);