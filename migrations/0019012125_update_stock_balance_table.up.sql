-- Add UUID primary key and change current_stock to decimal
ALTER TABLE stock_balances 
    ADD COLUMN uuid UUID DEFAULT uuid_generate_v4(),
    ALTER COLUMN current_stock TYPE DECIMAL(10,2);

-- Update the primary key constraint
ALTER TABLE stock_balances DROP CONSTRAINT IF EXISTS stock_balances_pkey;
ALTER TABLE stock_balances ADD PRIMARY KEY (uuid);