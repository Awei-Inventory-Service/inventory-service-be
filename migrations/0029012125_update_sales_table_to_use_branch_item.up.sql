-- Update sales table to use branch_products relationship
-- Remove old columns and add branch_product_id reference

ALTER TABLE sales 
DROP CONSTRAINT IF EXISTS fk_branch_id,
DROP COLUMN IF EXISTS branch_id,
DROP COLUMN IF EXISTS product_id;

-- Add branch_product_id column
ALTER TABLE sales 
ADD COLUMN branch_product_id UUID NOT NULL;

-- Add foreign key constraint to branch_products
ALTER TABLE sales 
ADD CONSTRAINT fk_sales_branch_product_id 
    FOREIGN KEY (branch_product_id) REFERENCES branch_products(uuid) 
    ON UPDATE CASCADE ON DELETE SET NULL;