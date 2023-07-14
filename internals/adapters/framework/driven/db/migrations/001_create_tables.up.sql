CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    curr_balance REAL NOT NULL, 
    operation VARCHAR NOT NULL
);