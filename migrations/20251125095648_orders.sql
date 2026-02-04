-- +goose Up
-- +goose StatementBegin

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_type 
        WHERE typname = 'order_status'
          AND typnamespace = (SELECT oid FROM pg_namespace WHERE nspname = current_schema())
    ) THEN
        CREATE TYPE order_status AS ENUM ('pending', 'paid', 'shipped', 'cancelled');
    END IF;
END $$;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price VARCHAR(30) DEFAULT '',
    status order_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_product_id ON orders(product_id);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
