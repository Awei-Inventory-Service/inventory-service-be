-- Update sales table quantity column to decimal for float support
ALTER TABLE sales 
ALTER COLUMN quantity TYPE DECIMAL(10,2);