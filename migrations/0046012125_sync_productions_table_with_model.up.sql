-- Migration: Sync productions table with Production model
-- This migration aligns the table structure with the Go model

-- Rename existing columns to match model expectations
ALTER TABLE productions RENAME COLUMN item_initial_amount TO initial_quantity;
ALTER TABLE productions RENAME COLUMN item_final_amount TO final_quantity;

-- Add missing columns from the model
ALTER TABLE productions ADD COLUMN waste_quantity DECIMAL(10,4) NOT NULL DEFAULT 0;
ALTER TABLE productions ADD COLUMN waste_percentage DECIMAL(10,4) NOT NULL DEFAULT 0;

-- Add missing unit columns
ALTER TABLE productions ADD COLUMN initial_unit VARCHAR(50) NOT NULL DEFAULT 'gram';
ALTER TABLE productions ADD COLUMN final_unit VARCHAR(50) NOT NULL DEFAULT 'gram';

-- Remove the default values after adding the columns
ALTER TABLE productions ALTER COLUMN waste_quantity DROP DEFAULT;
ALTER TABLE productions ALTER COLUMN waste_percentage DROP DEFAULT;
ALTER TABLE productions ALTER COLUMN initial_unit DROP DEFAULT;
ALTER TABLE productions ALTER COLUMN final_unit DROP DEFAULT;

-- Update constraints to use new column names
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_initial_amount_positive;
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_final_amount_positive;

-- Add new constraints with correct column names
ALTER TABLE productions ADD CONSTRAINT productions_initial_quantity_positive CHECK (initial_quantity > 0);
ALTER TABLE productions ADD CONSTRAINT productions_final_quantity_positive CHECK (final_quantity > 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_quantity_non_negative CHECK (waste_quantity >= 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_percentage_non_negative CHECK (waste_percentage >= 0);