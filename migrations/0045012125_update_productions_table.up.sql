-- Migration: Update productions table to match current model
-- Rename columns and add unit columns

-- Add new unit columns
ALTER TABLE productions 
ADD COLUMN initial_unit VARCHAR(50) NOT NULL DEFAULT 'units';

ALTER TABLE productions 
ADD COLUMN final_unit VARCHAR(50) NOT NULL DEFAULT 'units';

-- Rename existing columns
ALTER TABLE productions 
RENAME COLUMN item_initial_amount TO initial_quantity;

ALTER TABLE productions 
RENAME COLUMN item_final_amount TO final_quantity;

ALTER TABLE productions 
RENAME COLUMN waste TO waste_quantity;

-- Update constraint names to match new column names
ALTER TABLE productions DROP CONSTRAINT productions_initial_amount_positive;
ALTER TABLE productions DROP CONSTRAINT productions_final_amount_positive;
ALTER TABLE productions DROP CONSTRAINT productions_waste_non_negative;

-- Add updated constraints with new column names
ALTER TABLE productions ADD CONSTRAINT productions_initial_quantity_positive CHECK (initial_quantity > 0);
ALTER TABLE productions ADD CONSTRAINT productions_final_quantity_positive CHECK (final_quantity > 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_quantity_non_negative CHECK (waste_quantity >= 0);

-- Add constraints for unit columns
ALTER TABLE productions ADD CONSTRAINT productions_initial_unit_not_empty CHECK (initial_unit != '');
ALTER TABLE productions ADD CONSTRAINT productions_final_unit_not_empty CHECK (final_unit != '');