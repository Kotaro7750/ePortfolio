
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE ticker ADD COLUMN sector INT REFERENCES sector(id);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE ticker DROP COLUMN sector;
