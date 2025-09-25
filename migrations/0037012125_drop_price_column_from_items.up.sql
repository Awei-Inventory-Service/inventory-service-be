-- Migration: Drop price column from items table
-- Drop price constraint first, then drop the column

-- Drop the price positive constraint
ALTER TABLE items DROP CONSTRAINT IF EXISTS items_price_positive;

-- Drop the price column
ALTER TABLE items DROP COLUMN IF EXISTS price;