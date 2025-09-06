-- Create the item_category enum type
CREATE TYPE item_category AS ENUM (
    'processed',
    'half-processed', 
    'raw',
    'other'
);