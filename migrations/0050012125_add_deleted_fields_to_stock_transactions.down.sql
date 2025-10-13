-- Drop foreign key constraint first
ALTER TABLE stock_transactions
DROP CONSTRAINT IF EXISTS fk_deleted_by;

-- Drop the columns
ALTER TABLE stock_transactions 
DROP COLUMN IF EXISTS deleted_at,
DROP COLUMN IF EXISTS deleted_by;