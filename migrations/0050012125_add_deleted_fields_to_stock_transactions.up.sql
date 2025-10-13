ALTER TABLE stock_transactions 
ADD COLUMN deleted_at TIMESTAMP NULL,
ADD COLUMN deleted_by UUID NULL;

-- Add foreign key constraint for deleted_by
ALTER TABLE stock_transactions
ADD CONSTRAINT fk_deleted_by FOREIGN KEY (deleted_by) REFERENCES users(uuid)
ON UPDATE CASCADE
ON DELETE SET NULL;