
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE financial_report (
  ticker_id INT REFERENCES ticker(id),
  year INT NOT NULL,
  quarter INT NOT NULL,
  url TEXT NOT NULL,
  PRIMARY KEY(ticker_id,year,quarter)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS financial_report;
