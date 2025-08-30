CREATE TABLE item_compositions (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_item_id UUID NOT NULL,
    child_item_id UUID NOT NULL,
    ratio DECIMAL(10,4) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key constraints
    CONSTRAINT fk_item_compositions_parent_item 
        FOREIGN KEY (parent_item_id) REFERENCES items(uuid) 
        ON UPDATE CASCADE ON DELETE CASCADE,
    
    CONSTRAINT fk_item_compositions_child_item 
        FOREIGN KEY (child_item_id) REFERENCES items(uuid) 
        ON UPDATE CASCADE ON DELETE CASCADE,

    -- Business constraints
    CONSTRAINT item_compositions_ratio_positive CHECK (ratio > 0),
    
    -- Prevent self-referencing (item cannot be composed of itself)
    CONSTRAINT item_compositions_no_self_reference CHECK (parent_item_id != child_item_id),
    
    -- Prevent duplicate parent-child combinations
    CONSTRAINT item_compositions_unique_parent_child UNIQUE (parent_item_id, child_item_id)
);