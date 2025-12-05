-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_methods (
    id VARCHAR(2) PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    big_image_url VARCHAR(1024) NOT NULL,
    small_image_url VARCHAR(1024) NOT NULL,
    total_fee INT NOT NULL,
    category VARCHAR(255) NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_methods;
-- +goose StatementEnd
