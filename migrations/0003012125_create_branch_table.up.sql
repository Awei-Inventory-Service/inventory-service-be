CREATE TABLE branches (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    branch_manager_id UUID,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Add foreign key to user
    CONSTRAINT fk_branch_manager FOREIGN KEY (branch_manager_id) REFERENCES users(uuid)
);