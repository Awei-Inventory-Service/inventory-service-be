-- Migration: Create productions table
-- This table tracks production transformations of items

CREATE TABLE productions (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    source_item_id UUID NOT NULL,
    item_initial_amount DECIMAL(10,4) NOT NULL,
    final_item_id UUID NOT NULL,
    item_final_amount DECIMAL(10,4) NOT NULL,
    waste DECIMAL(10,4) NOT NULL,
    waste_percentage DECIMAL(10,4) NOT NULL,
    branch_id UUID NOT NULL,
    production_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for foreign key relationships
CREATE INDEX idx_productions_source_item_id ON productions(source_item_id);
CREATE INDEX idx_productions_final_item_id ON productions(final_item_id);
CREATE INDEX idx_productions_branch_id ON productions(branch_id);

-- Add foreign key constraints
ALTER TABLE productions 
    ADD CONSTRAINT fk_productions_branch 
    FOREIGN KEY (branch_id) REFERENCES branches(uuid) 
    ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE productions 
    ADD CONSTRAINT fk_productions_source_item 
    FOREIGN KEY (source_item_id) REFERENCES items(uuid) 
    ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE productions 
    ADD CONSTRAINT fk_productions_final_item 
    FOREIGN KEY (final_item_id) REFERENCES items(uuid) 
    ON UPDATE CASCADE ON DELETE CASCADE;

-- Add constraints for positive amounts
ALTER TABLE productions ADD CONSTRAINT productions_initial_amount_positive CHECK (item_initial_amount > 0);
ALTER TABLE productions ADD CONSTRAINT productions_final_amount_positive CHECK (item_final_amount > 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_non_negative CHECK (waste >= 0);
ALTER TABLE productions ADD CONSTRAINT productions_waste_percentage_non_negative CHECK (waste_percentage >= 0);