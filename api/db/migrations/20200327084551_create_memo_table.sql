
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE memo (
  id SERIAL NOT NULL,
  title TEXT NOT NULL,
  url TEXT NOT NULL,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS memo;
