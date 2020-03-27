
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE ticker ADD COLUMN dividened_month TEXT DEFAULT '3,6,9,12';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE ticker DROP COLUMN dividened_month;
