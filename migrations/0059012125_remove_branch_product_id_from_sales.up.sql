-- Remove branch_product_id column from sales table
-- Fix product_id type and set up proper foreign key constraints

-- Change product_id to UUID type to match products table
ALTER TABLE sales ALTER COLUMN product_id TYPE UUID USING product_id::UUID;

-- Make the new columns NOT NULL
ALTER TABLE sales ALTER COLUMN product_id SET NOT NULL;
ALTER TABLE sales ALTER COLUMN branch_id SET NOT NULL;

-- Add foreign key constraints for the new columns
ALTER TABLE sales ADD CONSTRAINT fk_sales_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL;

ALTER TABLE sales ADD CONSTRAINT fk_sales_product_id FOREIGN KEY (product_id) REFERENCES products(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL;

-- Drop the old branch_product_id column
ALTER TABLE sales DROP COLUMN branch_product_id;