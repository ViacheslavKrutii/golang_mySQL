DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'mydb') THEN
        CREATE DATABASE mydb;
    END IF;
END $$;

\c mydb;

CREATE TABLE IF NOT EXISTS orders (
    orderID SERIAL PRIMARY KEY,
    Username VARCHAR(255),
    order_body JSON
);
