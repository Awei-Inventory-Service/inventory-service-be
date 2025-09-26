-- Migration: Create product_recipes table to replace product_compositions
CREATE TABLE product_recipes (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    item_id UUID NOT NULL,
    ratio DECIMAL(10,4) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key constraints
    CONSTRAINT fk_product_recipes_product 
        FOREIGN KEY (product_id) REFERENCES products(uuid) 
        ON UPDATE CASCADE ON DELETE CASCADE,
    
    CONSTRAINT fk_product_recipes_item 
        FOREIGN KEY (item_id) REFERENCES items(uuid) 
        ON UPDATE CASCADE ON DELETE CASCADE,

    -- Business constraints
    CONSTRAINT product_recipes_ratio_positive CHECK (ratio > 0),
    
    -- Prevent duplicate product-item combinations
    CONSTRAINT product_recipes_unique_product_item UNIQUE (product_id, item_id)
);