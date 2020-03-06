
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE ticker (
  id SERIAL NOT NULL,
  ticker VARCHAR(255) NOT NULL,
  dividened REAL NOT NULL,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS ticker;
