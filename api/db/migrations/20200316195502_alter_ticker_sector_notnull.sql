
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE ticker ALTER COLUMN sector SET NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE ticker ALTER COLUMN sector DROP NOT NULL;
