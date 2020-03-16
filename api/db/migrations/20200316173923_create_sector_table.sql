
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE sector (
  id SERIAL NOT NULL,
  sector VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);

INSERT INTO sector (sector) VALUES ('Information Technology');
INSERT INTO sector (sector) VALUES ('Health Care');
INSERT INTO sector (sector) VALUES ('Financials');
INSERT INTO sector (sector) VALUES ('Communication Services');
INSERT INTO sector (sector) VALUES ('Consumer Discretionary');
INSERT INTO sector (sector) VALUES ('Industrials');
INSERT INTO sector (sector) VALUES ('Consumer Staples');
INSERT INTO sector (sector) VALUES ('Energy');
INSERT INTO sector (sector) VALUES ('Utilities');
INSERT INTO sector (sector) VALUES ('Real Estate');
INSERT INTO sector (sector) VALUES ('Materials');


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS sector;
