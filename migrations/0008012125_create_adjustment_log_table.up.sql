CREATE TABLE adjustment_logs (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    branch_id UUID NOT NULL,
    item_id UUID NOT NULL,
    previous_stock INT NOT NULL,
    new_stock INT NOT NULL,
    adjustor_id UUID NOT NULL,
    remarks TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Foreign key to branch_id
    CONSTRAINT fk_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Foreign key to item_id
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES items(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Foreign key to adjustor_id
    CONSTRAINT fk_adjustor_id FOREIGN KEY (adjustor_id) REFERENCES users(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);