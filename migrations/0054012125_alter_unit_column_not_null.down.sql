-- Revert unit column to nullable
ALTER TABLE inventory_transfer_items ALTER COLUMN unit DROP NOT NULL;
ALTER TABLE inventory_transfer_items ALTER COLUMN unit DROP DEFAULT;