-- Migration: Make waste_quantity and waste_percentage columns nullable in production_items table
-- This allows production items to have optional waste tracking

-- Remove NOT NULL constraints and set default NULL
ALTER TABLE production_items ALTER COLUMN waste_quantity DROP NOT NULL;
ALTER TABLE production_items ALTER COLUMN waste_percentage DROP NOT NULL;

-- Remove existing default values
ALTER TABLE production_items ALTER COLUMN waste_quantity DROP DEFAULT;
ALTER TABLE production_items ALTER COLUMN waste_percentage DROP DEFAULT;

-- Update existing records with 0 values to NULL (optional - remove if you want to keep existing 0 values)
UPDATE production_items SET waste_quantity = NULL WHERE waste_quantity = 0;
UPDATE production_items SET waste_percentage = NULL WHERE waste_percentage = 0;