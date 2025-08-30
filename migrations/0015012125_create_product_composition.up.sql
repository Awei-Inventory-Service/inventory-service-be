CREATE TABLE product_compositions (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    item_id UUID NOT NULL,
    ratio DECIMAL(10,4) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key constraints
    CONSTRAINT fk_product_compositions_product 
        FOREIGN KEY (product_id) REFERENCES products(uuid) 
        ON UPDATE CASCADE ON DELETE CASCADE,
    
    CONSTRAINT fk_product_compositions_item 
        FOREIGN KEY (item_id) REFERENCES items(uuid) 
        ON UPDATE CASCADE ON DELETE CASCADE,

    -- Business constraints
    CONSTRAINT product_compositions_ratio_positive CHECK (ratio > 0),
    
    -- Prevent duplicate product-item combinations
    CONSTRAINT product_compositions_unique_product_item UNIQUE (product_id, item_id)
);
