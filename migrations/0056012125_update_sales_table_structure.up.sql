-- Update sales table to match new model structure
-- Add branch_id and product_id columns to replace branch_product_id

-- Add the new columns
ALTER TABLE sales ADD COLUMN IF NOT EXISTS branch_id UUID;
ALTER TABLE sales ADD COLUMN IF NOT EXISTS product_id VARCHAR(255);

-- Note: price and transaction_date already exist, so we don't need to add them
-- Just ensure they have the right constraints

-- Populate the new columns based on existing branch_product_id
-- This is a data migration step - you may need to customize this based on your data
-- For now, we'll add the columns and they can be populated manually or with additional logic

-- Drop the old foreign key constraint
ALTER TABLE sales DROP CONSTRAINT IF EXISTS fk_sales_branch_product_id;

-- Make the new columns NOT NULL (after populating them)
-- Note: You'll need to populate these columns with actual data before setting NOT NULL
-- ALTER TABLE sales ALTER COLUMN branch_id SET NOT NULL;
-- ALTER TABLE sales ALTER COLUMN product_id SET NOT NULL;

-- Add foreign key constraints for the new columns
-- Note: Uncomment these after populating the columns
-- ALTER TABLE sales ADD CONSTRAINT fk_sales_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid)
--     ON UPDATE CASCADE
--     ON DELETE SET NULL;

-- ALTER TABLE sales ADD CONSTRAINT fk_sales_product_id FOREIGN KEY (product_id) REFERENCES products(uuid)
--     ON UPDATE CASCADE
--     ON DELETE SET NULL;

-- Remove the old branch_product_id column (after data migration)
-- ALTER TABLE sales DROP COLUMN branch_product_id;