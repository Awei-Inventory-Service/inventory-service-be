-- Add back the columns that were removed
ALTER TABLE sales ADD COLUMN product_id uuid;
ALTER TABLE sales ADD COLUMN type varchar(255) NOT NULL DEFAULT '';
ALTER TABLE sales ADD COLUMN quantity decimal NOT NULL DEFAULT 0;
ALTER TABLE sales ADD COLUMN cost decimal NOT NULL DEFAULT 0;
ALTER TABLE sales ADD COLUMN price decimal NOT NULL DEFAULT 0;