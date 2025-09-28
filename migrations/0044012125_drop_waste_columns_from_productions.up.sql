-- Migration: Drop waste and waste_percentage columns from productions table
-- These columns are no longer needed

-- Drop the constraints first
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_waste_non_negative;
ALTER TABLE productions DROP CONSTRAINT IF EXISTS productions_waste_percentage_non_negative;

-- Drop the columns
ALTER TABLE productions DROP COLUMN IF EXISTS waste;
ALTER TABLE productions DROP COLUMN IF EXISTS waste_percentage;