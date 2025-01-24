CREATE TABLE stock_balances (
    branch_id UUID NOT NULL,
    item_id UUID NOT NULL,
    current_stock INT NOT NULL,
    -- Add foreign key to branch
    CONSTRAINT fk_branch_id FOREIGN KEY (branch_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Add foreign key to item
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES items(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    PRIMARY KEY (branch_id, item_id)
);