-- Migration rollback: Restore NOT NULL constraints and default values for waste columns

-- Set NULL values back to 0 before adding NOT NULL constraint
UPDATE production_items SET waste_quantity = 0 WHERE waste_quantity IS NULL;
UPDATE production_items SET waste_percentage = 0 WHERE waste_percentage IS NULL;

-- Restore default values
ALTER TABLE production_items ALTER COLUMN waste_quantity SET DEFAULT 0;
ALTER TABLE production_items ALTER COLUMN waste_percentage SET DEFAULT 0;

-- Restore NOT NULL constraints
ALTER TABLE production_items ALTER COLUMN waste_quantity SET NOT NULL;
ALTER TABLE production_items ALTER COLUMN waste_percentage SET NOT NULL;