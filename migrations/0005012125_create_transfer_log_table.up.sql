CREATE TABLE transfer_logs (
    branch_origin_id UUID NOT NULL,
    branch_dest_id UUID NOT NULL,
    item_id UUID NOT NULL,
    issuer_id UUID NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Add foreign key to branch_origin_id
    CONSTRAINT fk_branch_origin_id FOREIGN KEY (branch_origin_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Add foreign key to branch_dest_id
    CONSTRAINT fk_branch_dest_id FOREIGN KEY (branch_dest_id) REFERENCES branches(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Add foreign key to item_id
    CONSTRAINT fk_item_id FOREIGN KEY (item_id) REFERENCES items(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL,
    -- Add foreign key to issuer_id
    CONSTRAINT fk_issuer_id FOREIGN KEY (issuer_id) REFERENCES users(uuid)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);