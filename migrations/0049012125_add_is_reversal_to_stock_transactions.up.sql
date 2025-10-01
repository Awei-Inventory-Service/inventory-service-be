-- Migration: Add is_reversal column to stock_transactions table
-- This column tracks whether a stock transaction is a reversal of another transaction

ALTER TABLE stock_transactions ADD COLUMN is_reversal BOOLEAN NULL;