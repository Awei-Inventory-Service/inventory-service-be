-- Drop stock_balances table as it's redundant with branch_items table
-- branch_items provides the same functionality with additional features:
-- - price tracking
-- - proper foreign key constraints with CASCADE
-- - timestamps (created_at, updated_at)
-- - better indexing
-- - unique constraint on (branch_id, item_id)

DROP TABLE IF EXISTS stock_balances;