-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE ticker ADD CONSTRAINT ticker_ticker_unique unique(ticker);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

