-- Create inventory_transfers table
CREATE TABLE inventory_transfers (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_origin_id UUID NOT NULL,
    branch_destination_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL,
    transfer_date TIMESTAMP NOT NULL,
    remarks TEXT,
    issuer_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraints
    CONSTRAINT fk_inventory_transfers_branch_origin 
        FOREIGN KEY (branch_origin_id) REFERENCES branches(uuid) 
        ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_inventory_transfers_branch_destination 
        FOREIGN KEY (branch_destination_id) REFERENCES branches(uuid) 
        ON UPDATE CASCADE ON DELETE SET NULL
);

-- Create indexes
CREATE INDEX idx_inventory_transfers_branch_origin_id ON inventory_transfers(branch_origin_id);
CREATE INDEX idx_inventory_transfers_branch_destination_id ON inventory_transfers(branch_destination_id);
CREATE INDEX idx_inventory_transfers_status ON inventory_transfers(status);
CREATE INDEX idx_inventory_transfers_transfer_date ON inventory_transfers(transfer_date);