-- +goose Up
-- +goose StatementBegin
INSERT INTO payment_methods (id, name, small_image_url, big_image_url, total_fee, category)
VALUES ('QD',
        'Dompet Kita QRIS',
        'https://www.google.com',
        'https://www.google.com',
        0,
        'QRIS'
    );
INSERT INTO payment_methods (id, name, small_image_url, big_image_url, total_fee, category)
VALUES ('DK',
        'Dompet Kita',
        'https://www.google.com',
        'https://www.google.com',
        1000,
        'EWALLET'
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM payment_methods;
-- +goose StatementEnd
