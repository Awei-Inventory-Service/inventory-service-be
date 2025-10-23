-- Add unit column to inventory_transfer_items table
ALTER TABLE inventory_transfer_items ADD COLUMN unit VARCHAR(50) NOT NULL DEFAULT '';