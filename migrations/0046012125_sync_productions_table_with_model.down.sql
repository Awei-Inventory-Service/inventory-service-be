-- Migration rollback: Revert productions table sync
-- This reverts the table structure back to the original state

-- Remove new constraints
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_initial_quantity_positive;
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_final_quantity_positive;
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_waste_quantity_non_negative;
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_waste_percentage_non_negative;

-- Remove added columns
ALTER TABLE productions DROP COLUMN IF EXISTS waste_quantity;
ALTER TABLE productions DROP COLUMN IF EXISTS waste_percentage;
ALTER TABLE productions DROP COLUMN IF EXISTS initial_unit;
ALTER TABLE productions DROP COLUMN IF EXISTS final_unit;

-- Rename columns back to original names
ALTER TABLE productions RENAME COLUMN initial_quantity TO item_initial_amount;
ALTER TABLE productions RENAME COLUMN final_quantity TO item_final_amount;

-- Restore original constraints
ALTER TABLE productions ADD CONSTRAINT productions_initial_amount_positive CHECK (item_initial_amount > 0);
ALTER TABLE productions ADD CONSTRAINT productions_final_amount_positive CHECK (item_final_amount > 0);