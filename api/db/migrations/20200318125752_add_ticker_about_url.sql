
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE ticker ADD COLUMN about_url TEXT DEFAULT '';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE ticker DROP COLUMN about_url;
