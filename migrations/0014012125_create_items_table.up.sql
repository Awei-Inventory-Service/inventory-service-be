CREATE TABLE items (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    price DECIMAL(10,2),  -- Note: No NOT NULL to match GORM tag exactly
    unit VARCHAR(255) NOT NULL,
    portion_size DECIMAL(10,4) DEFAULT 1,
    supplier_id UUID,  -- Made nullable to match onDelete:SET NULL constraint
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraint matching GORM specification
    CONSTRAINT fk_items_supplier 
        FOREIGN KEY (supplier_id) REFERENCES suppliers(uuid) 
        ON UPDATE CASCADE ON DELETE SET NULL,
    
    -- Business logic constraints (based on validation tags)
    CONSTRAINT items_price_positive CHECK (price IS NULL OR price > 0),
    CONSTRAINT items_portion_size_positive CHECK (portion_size > 0),
    CONSTRAINT items_name_not_empty CHECK (TRIM(name) != ''),
    CONSTRAINT items_unit_not_empty CHECK (TRIM(unit) != '')
);
