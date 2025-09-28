-- Migration rollback: Revert productions table to single-source structure
-- This migration reverts the multi-source production changes

-- Step 1: Add back single-source columns to productions table
ALTER TABLE productions ADD COLUMN source_item_id UUID;
ALTER TABLE productions ADD COLUMN initial_quantity DECIMAL(10,4);
ALTER TABLE productions ADD COLUMN initial_unit VARCHAR(50);
ALTER TABLE productions ADD COLUMN waste_quantity DECIMAL(10,4) DEFAULT 0;
ALTER TABLE productions ADD COLUMN waste_percentage DECIMAL(10,4) DEFAULT 0;

-- Step 2: Migrate first source item from production_items back to productions
-- This only takes the first source item if multiple exist
WITH first_source AS (
    SELECT DISTINCT ON (production_id) 
        production_id,
        source_item_id,
        quantity,
        unit,
        waste_quantity,
        waste_percentage
    FROM production_items 
    ORDER BY production_id, created_at
)
UPDATE productions 
SET 
    source_item_id = first_source.source_item_id,
    initial_quantity = first_source.quantity,
    initial_unit = first_source.unit,
    waste_quantity = first_source.waste_quantity,
    waste_percentage = first_source.waste_percentage
FROM first_source 
WHERE productions.uuid = first_source.production_id;

-- Step 3: Remove default values
ALTER TABLE productions ALTER COLUMN waste_quantity DROP DEFAULT;
ALTER TABLE productions ALTER COLUMN waste_percentage DROP DEFAULT;

-- Step 4: Add back constraints and indexes
CREATE INDEX IF NOT EXISTS idx_productions_source_item_id ON productions(source_item_id);

ALTER TABLE productions 
    ADD CONSTRAINT fk_productions_source_item 
    FOREIGN KEY (source_item_id) REFERENCES items(uuid) 
    ON UPDATE CASCADE ON DELETE CASCADE;

-- Step 5: Add back constraints for positive values
ALTER TABLE productions ADD CONSTRAINT productions_initial_quantity_positive CHECK (initial_quantity > 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_quantity_non_negative CHECK (waste_quantity >= 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_percentage_non_negative CHECK (waste_percentage >= 0);

-- Step 6: Drop production_items table
DROP TABLE IF EXISTS production_items;