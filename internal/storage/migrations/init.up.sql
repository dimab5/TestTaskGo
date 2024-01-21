CREATE TABLE IF NOT EXISTS Wallets(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    balance DECIMAL(10, 2) DEFAULT 100.0
);

CREATE TABLE IF NOT EXISTS Transaction_History(
    from_id UUID NOT NULL,
    to_id UUID NOT NULL,
    amount DECIMAL(10, 2),
    time TIMESTAMP,
    response_status INTEGER
)