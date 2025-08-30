-- Create products table
CREATE TABLE products (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    unit VARCHAR(255) NOT NULL,
    selling_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Business constraints
    CONSTRAINT products_selling_price_positive CHECK (selling_price > 0),
    CONSTRAINT products_name_not_empty CHECK (TRIM(name) != ''),
    CONSTRAINT products_code_not_empty CHECK (TRIM(code) != ''),
    CONSTRAINT products_category_not_empty CHECK (TRIM(category) != ''),
    CONSTRAINT products_unit_not_empty CHECK (TRIM(unit) != ''),
    
    -- Optional: Unique constraint on code if it should be unique
    CONSTRAINT products_code_unique UNIQUE (code)
);

