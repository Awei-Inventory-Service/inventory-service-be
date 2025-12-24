CREATE TABLE sales_products (
    uuid uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    sales_id uuid NOT NULL REFERENCES sales(uuid) ON UPDATE CASCADE ON DELETE CASCADE,
    product_id uuid NOT NULL,
    quantity decimal NOT NULL,
    type varchar(255) NOT NULL,
    cost decimal NOT NULL,
    price decimal NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);