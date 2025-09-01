ALTER TABLE stock_transactions 
ALTER COLUMN quantity TYPE DECIMAL USING quantity::DECIMAL;