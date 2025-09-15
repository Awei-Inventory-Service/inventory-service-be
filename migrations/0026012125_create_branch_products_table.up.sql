-- Create branch_products table
CREATE TABLE branch_products (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_id UUID NOT NULL,
    product_id UUID NOT NULL,
    stock DECIMAL(10,2),
    buy_price DECIMAL(10,2),
    selling_price DECIMAL(10,2),
    supplier_id UUID,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key constraints
    CONSTRAINT fk_branch_products_branch_id 
        FOREIGN KEY (branch_id) REFERENCES branches(uuid) 
        ON UPDATE CASCADE ON DELETE SET NULL,
    
    CONSTRAINT fk_branch_products_product_id 
        FOREIGN KEY (product_id) REFERENCES products(uuid) 
        ON UPDATE CASCADE ON DELETE SET NULL,
    
    CONSTRAINT fk_branch_products_supplier_id 
        FOREIGN KEY (supplier_id) REFERENCES suppliers(uuid) 
        ON UPDATE CASCADE ON DELETE SET NULL,

    -- Business constraints
    CONSTRAINT branch_products_stock_non_negative CHECK (stock IS NULL OR stock >= 0),
    CONSTRAINT branch_products_buy_price_positive CHECK (buy_price IS NULL OR buy_price > 0),
    CONSTRAINT branch_products_selling_price_positive CHECK (selling_price IS NULL OR selling_price > 0),

    -- Unique constraint - one record per branch-product combination
    CONSTRAINT branch_products_branch_product_unique UNIQUE (branch_id, product_id)
);