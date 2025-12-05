-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    payment_method_id VARCHAR(2) NOT NULL REFERENCES payment_methods(id) ON UPDATE CASCADE ON DELETE CASCADE,
    amount BIGINT NOT NULL,
    status VARCHAR(255) NOT NULL DEFAULT 'PENDING',
    expiry_time TIMESTAMP NOT NULL,
    callback_url VARCHAR(1024) NOT NULL,
    redirect_url VARCHAR(1024) NOT NULL,
    merchant_id VARCHAR(255) NOT NULL,
    merchant_order_id VARCHAR(255) NOT NULL,
    customer_email VARCHAR(255) NOT NULL,
    customer_name VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(20) NOT NULL,
    product_details VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
