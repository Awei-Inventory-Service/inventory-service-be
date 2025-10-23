-- Make unit column NOT NULL with default empty string
UPDATE inventory_transfer_items SET unit = '' WHERE unit IS NULL;
ALTER TABLE inventory_transfer_items ALTER COLUMN unit SET NOT NULL;
ALTER TABLE inventory_transfer_items ALTER COLUMN unit SET DEFAULT '';