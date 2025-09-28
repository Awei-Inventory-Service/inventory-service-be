-- Rollback migration: Revert productions table changes

-- Remove unit columns
ALTER TABLE productions DROP COLUMN initial_unit;
ALTER TABLE productions DROP COLUMN final_unit;

-- Rename columns back to original names
ALTER TABLE productions 
RENAME COLUMN initial_quantity TO item_initial_amount;

ALTER TABLE productions 
RENAME COLUMN final_quantity TO item_final_amount;

ALTER TABLE productions 
RENAME COLUMN waste_quantity TO waste;

-- Remove new constraints
ALTER TABLE productions DROP CONSTRAINT productions_initial_quantity_positive;
ALTER TABLE productions DROP CONSTRAINT productions_final_quantity_positive;
ALTER TABLE productions DROP CONSTRAINT productions_waste_quantity_non_negative;
ALTER TABLE productions DROP CONSTRAINT productions_initial_unit_not_empty;
ALTER TABLE productions DROP CONSTRAINT productions_final_unit_not_empty;

-- Restore original constraints
ALTER TABLE productions ADD CONSTRAINT productions_initial_amount_positive CHECK (item_initial_amount > 0);
ALTER TABLE productions ADD CONSTRAINT productions_final_amount_positive CHECK (item_final_amount > 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_non_negative CHECK (waste >= 0);