CREATE TYPE invoice_status AS ENUM ('paid', 'unpaid');
CREATE TABLE invoices (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    file_url TEXT NOT NULL,
    amount DECIMAL NOT NULL CHECK (amount > 0),
    amount_owed DECIMAL NOT NULL DEFAULT 0 CHECK (amount_owed >= 0),
    status invoice_status NOT NULL DEFAULT 'unpaid',
    notes TEXT,
    invoice_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT status_amount_owed_check CHECK (
        (status = 'unpaid' AND amount_owed = amount) OR
        (status = 'paid')
    ),

    CONSTRAINT amount_owed_check CHECK (amount_owed <= amount)
);