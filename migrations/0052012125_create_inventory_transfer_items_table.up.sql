-- Create inventory_transfer_items table
CREATE TABLE inventory_transfer_items (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    inventory_transfer_id UUID NOT NULL,
    item_id UUID NOT NULL,
    item_quantity DECIMAL(15,4) NOT NULL,
    item_cost DECIMAL(15,4) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraints
    CONSTRAINT fk_inventory_transfer_items_transfer 
        FOREIGN KEY (inventory_transfer_id) REFERENCES inventory_transfers(uuid) 
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_inventory_transfer_items_item 
        FOREIGN KEY (item_id) REFERENCES items(uuid) 
        ON UPDATE CASCADE ON DELETE SET NULL
);

-- Create indexes
CREATE INDEX idx_inventory_transfer_items_transfer_id ON inventory_transfer_items(inventory_transfer_id);
CREATE INDEX idx_inventory_transfer_items_item_id ON inventory_transfer_items(item_id);