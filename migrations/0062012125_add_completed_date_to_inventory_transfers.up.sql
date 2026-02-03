-- Add completed_date column to inventory_transfers table
ALTER TABLE inventory_transfers ADD COLUMN completed_date TIMESTAMP NULL;

-- Create index for completed_date for better query performance
CREATE INDEX idx_inventory_transfers_completed_date ON inventory_transfers(completed_date);
