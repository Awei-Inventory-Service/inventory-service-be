-- Restore branch_product_id column to sales table
-- Remove new foreign key constraints and add back branch_product_id

-- Drop the new foreign key constraints
ALTER TABLE sales DROP CONSTRAINT IF EXISTS fk_sales_branch_id;
ALTER TABLE sales DROP CONSTRAINT IF EXISTS fk_sales_product_id;

-- Add back branch_product_id column
ALTER TABLE sales ADD COLUMN branch_product_id UUID NOT NULL;

-- Make the new columns nullable again
ALTER TABLE sales ALTER COLUMN product_id DROP NOT NULL;
ALTER TABLE sales ALTER COLUMN branch_id DROP NOT NULL;