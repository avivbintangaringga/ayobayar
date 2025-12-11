-- +goose Up
-- +goose StatementBegin
INSERT INTO payments (
    id,
    payment_method_id,
    amount,
    status,
    expiry_time,
    callback_url,
    redirect_url,
    merchant_id,
    merchant_order_id,
    customer_email,
    customer_name,
    customer_phone,
    product_details
) VALUES (
    'testid',
    'QD',
    153000,
    'PENDING',
    NOW() + INTERVAL '1 year',
    'https://www.google.com',
    'https://www.google.com',
    'SDK12121212',
    'QDQD1010101002',
    'mail@test.com',
    'Dummy User',
    '1234567890',
    'Dummy Product'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM payments;
-- +goose StatementEnd
