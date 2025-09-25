-- Add transaction_date column to sales table with default current timestamp
ALTER TABLE sales 
ADD COLUMN transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;