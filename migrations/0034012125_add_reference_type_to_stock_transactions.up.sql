-- Add reference_type column to stock_transactions table
ALTER TABLE stock_transactions 
ADD COLUMN reference_type VARCHAR(100);