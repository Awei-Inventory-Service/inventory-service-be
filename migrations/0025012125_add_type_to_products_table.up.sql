-- Add type column to products table
ALTER TABLE products 
ADD COLUMN type product_type NOT NULL DEFAULT 'produced';

-- Add constraint to ensure type is valid
-- (The enum already handles this, but explicit constraint for clarity)
ALTER TABLE products 
ADD CONSTRAINT products_type_not_null CHECK (type IS NOT NULL);