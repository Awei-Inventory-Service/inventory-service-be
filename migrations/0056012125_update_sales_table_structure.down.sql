-- Revert sales table structure changes
-- This will restore the table to use branch_product_id instead of separate product_id and branch_id

-- Drop the new foreign key constraints
ALTER TABLE sales DROP CONSTRAINT IF EXISTS fk_sales_product_id;
ALTER TABLE sales DROP CONSTRAINT IF EXISTS fk_sales_branch_id;

-- Add back branch_product_id column
ALTER TABLE sales ADD COLUMN IF NOT EXISTS branch_product_id VARCHAR(255);

-- Drop the new columns
ALTER TABLE sales DROP COLUMN IF EXISTS product_id;
ALTER TABLE sales DROP COLUMN IF EXISTS price;
ALTER TABLE sales DROP COLUMN IF EXISTS transaction_date;

-- Change quantity back to INT
ALTER TABLE sales ALTER COLUMN quantity TYPE INT;

-- Add foreign key constraint for branch_product_id (if branch_products table exists)
-- ALTER TABLE sales ADD CONSTRAINT fk_branch_product_id FOREIGN KEY (branch_product_id) REFERENCES branch_products(id)
--     ON UPDATE CASCADE
--     ON DELETE SET NULL;