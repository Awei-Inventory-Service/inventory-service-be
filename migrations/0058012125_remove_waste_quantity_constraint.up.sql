-- Remove the non-negative constraint on waste_quantity
-- This allows negative waste quantity values in production_items

ALTER TABLE production_items DROP CONSTRAINT IF EXISTS production_items_waste_quantity_non_negative;