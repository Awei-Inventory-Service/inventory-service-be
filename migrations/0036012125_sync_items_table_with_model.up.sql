-- Migration: Sync items table with model.Item structure
-- Drop portion_size column and related constraint

-- Drop the constraint first
ALTER TABLE items DROP CONSTRAINT IF EXISTS items_portion_size_positive;

-- Drop the portion_size column
ALTER TABLE items DROP COLUMN IF EXISTS portion_size;