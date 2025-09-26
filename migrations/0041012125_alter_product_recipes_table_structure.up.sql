-- Migration: Alter product_recipes table to match ProductRecipe model
-- Replace ratio column with amount and add unit column

-- Drop the ratio constraint first
ALTER TABLE product_recipes DROP CONSTRAINT IF EXISTS product_recipes_ratio_positive;

-- Drop ratio column and add amount and unit columns
ALTER TABLE product_recipes DROP COLUMN IF EXISTS ratio;
ALTER TABLE product_recipes ADD COLUMN amount DECIMAL(10,4) NOT NULL DEFAULT 0;
ALTER TABLE product_recipes ADD COLUMN unit TEXT;

-- Add constraint for positive amount
ALTER TABLE product_recipes ADD CONSTRAINT product_recipes_amount_positive CHECK (amount > 0);