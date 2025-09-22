-- Rename cost column to price in sales table to better represent selling price
ALTER TABLE sales 
RENAME COLUMN cost TO price;