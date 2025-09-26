-- Migration: Create inventories table
CREATE TABLE inventories (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_id UUID NOT NULL,
    item_id UUID NOT NULL,
    stock DECIMAL(10,2) NOT NULL,
    value DECIMAL(10,2),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key constraints
    CONSTRAINT fk_inventories_branch
        FOREIGN KEY (branch_id) REFERENCES branches(uuid)
        ON UPDATE CASCADE ON DELETE SET NULL,
    
    CONSTRAINT fk_inventories_item
        FOREIGN KEY (item_id) REFERENCES items(uuid)
        ON UPDATE CASCADE ON DELETE SET NULL,

    -- Business constraints
    CONSTRAINT inventories_stock_non_negative CHECK (stock >= 0),
    CONSTRAINT inventories_value_non_negative CHECK (value IS NULL OR value >= 0),
    
    -- Unique constraint to prevent duplicate branch-item combinations
    CONSTRAINT inventories_unique_branch_item UNIQUE (branch_id, item_id)
);

-- Create indexes for better query performance
CREATE INDEX idx_inventories_branch_id ON inventories(branch_id);
CREATE INDEX idx_inventories_item_id ON inventories(item_id);
CREATE INDEX idx_inventories_stock ON inventories(stock);