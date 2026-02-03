-- Remove index
DROP INDEX IF EXISTS idx_inventory_transfers_completed_date;

-- Remove completed_date column from inventory_transfers table
ALTER TABLE inventory_transfers DROP COLUMN completed_date;
