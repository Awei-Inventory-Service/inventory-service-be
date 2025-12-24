-- Re-add the non-negative constraint on waste_quantity
-- This restores the original constraint that prevents negative waste quantity values

ALTER TABLE production_items ADD CONSTRAINT production_items_waste_quantity_non_negative 
    CHECK (waste_quantity >= 0::numeric);