-- Add purchase_date column to purchases table
-- This represents the actual date when the purchase was made (business date)
-- Separate from created_at which is the system timestamp when the record was created

ALTER TABLE purchases 
ADD COLUMN purchase_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- Update existing records to use created_at as purchase_date initially
UPDATE purchases SET purchase_date = created_at WHERE purchase_date IS NULL;

-- Add comment for clarity
COMMENT ON COLUMN purchases.purchase_date IS 'Business date when the purchase was actually made';