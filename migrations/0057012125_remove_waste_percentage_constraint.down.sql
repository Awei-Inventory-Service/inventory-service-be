-- Re-add the non-negative constraint on waste_percentage
-- This restores the original constraint that prevents negative waste percentage values

ALTER TABLE production_items ADD CONSTRAINT production_items_waste_percentage_non_negative 
    CHECK (waste_percentage >= 0::numeric);