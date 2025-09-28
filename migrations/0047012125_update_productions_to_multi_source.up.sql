-- Migration: Update productions table to support multiple source items
-- This migration restructures the production system to handle multiple ingredients

-- Step 1: Create production_items table for junction
CREATE TABLE production_items (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    production_id UUID NOT NULL,
    source_item_id UUID NOT NULL,
    quantity DECIMAL(10,4) NOT NULL,
    unit VARCHAR(50) NOT NULL,
    waste_quantity DECIMAL(10,4) NOT NULL DEFAULT 0,
    waste_percentage DECIMAL(10,4) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Step 2: Create indexes for performance
CREATE INDEX idx_production_items_production_id ON production_items(production_id);
CREATE INDEX idx_production_items_source_item_id ON production_items(source_item_id);

-- Step 3: Add foreign key constraints
ALTER TABLE production_items 
    ADD CONSTRAINT fk_production_items_production 
    FOREIGN KEY (production_id) REFERENCES productions(uuid) 
    ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE production_items 
    ADD CONSTRAINT fk_production_items_source_item 
    FOREIGN KEY (source_item_id) REFERENCES items(uuid) 
    ON UPDATE CASCADE ON DELETE CASCADE;

-- Step 4: Migrate existing single-source production data to new structure
-- This handles any existing productions by creating corresponding production_items records
INSERT INTO production_items (production_id, source_item_id, quantity, unit, waste_quantity, waste_percentage, created_at, updated_at)
SELECT 
    uuid as production_id,
    source_item_id,
    initial_quantity as quantity,
    initial_unit as unit,
    waste_quantity,
    waste_percentage,
    created_at,
    updated_at
FROM productions 
WHERE source_item_id IS NOT NULL;

-- Step 5: Remove old single-source columns from productions table
ALTER TABLE productions DROP COLUMN IF EXISTS source_item_id;
ALTER TABLE productions DROP COLUMN IF EXISTS initial_quantity;
ALTER TABLE productions DROP COLUMN IF EXISTS initial_unit;
ALTER TABLE productions DROP COLUMN IF EXISTS waste_quantity;
ALTER TABLE productions DROP COLUMN IF EXISTS waste_percentage;

-- Step 6: Add constraints for positive values
ALTER TABLE production_items ADD CONSTRAINT production_items_quantity_positive CHECK (quantity > 0);
ALTER TABLE production_items ADD CONSTRAINT production_items_waste_quantity_non_negative CHECK (waste_quantity >= 0);
ALTER TABLE production_items ADD CONSTRAINT production_items_waste_percentage_non_negative CHECK (waste_percentage >= 0);