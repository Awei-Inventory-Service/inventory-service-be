-- Add cost column back to sales table for profit tracking
-- Now we'll have both cost (purchase/base cost) and price (selling price)
ALTER TABLE sales 
ADD COLUMN cost DECIMAL(10,2) NOT NULL DEFAULT 0;