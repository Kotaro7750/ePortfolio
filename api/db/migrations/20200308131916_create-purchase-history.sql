
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE purchase_history (
  id SERIAL NOT NULL,
  ticker_id INT REFERENCES ticker(id),
  date DATE NOT NULL,
  share INT NOT NULL,
  cost REAL NOT NULL,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS purchase_history;
