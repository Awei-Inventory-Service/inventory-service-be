-- Migration rollback: Remove is_reversal column from stock_transactions table

ALTER TABLE stock_transactions DROP COLUMN IF EXISTS is_reversal;