-- Remove the non-negative constraint on waste_percentage
-- This allows negative waste percentage values in production_items

ALTER TABLE production_items DROP CONSTRAINT IF EXISTS production_items_waste_percentage_non_negative;