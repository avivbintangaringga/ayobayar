-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    payment_method_id VARCHAR(2) NOT NULL REFERENCES payment_methods(id),
    amount BIGINT NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    callback_url VARCHAR(1024) NOT NULL,
    merchant_id VARCHAR(255) NOT NULL,
    merchant_order_id VARCHAR(255) NOT NULL,
    user_email VARCHAR(255) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
