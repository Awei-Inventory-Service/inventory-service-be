-- Alter the items table to use item_category enum
ALTER TABLE items 
ALTER COLUMN category TYPE item_category 
USING category::item_category;