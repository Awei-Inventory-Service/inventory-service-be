-- Remove columns that will move to sales_products table
ALTER TABLE sales DROP COLUMN product_id;
ALTER TABLE sales DROP COLUMN type;
ALTER TABLE sales DROP COLUMN quantity;
ALTER TABLE sales DROP COLUMN cost;
ALTER TABLE sales DROP COLUMN price;