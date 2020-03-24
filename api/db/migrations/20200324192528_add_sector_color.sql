
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE sector ADD COLUMN color TEXT;

UPDATE sector SET color = '#fc0121' WHERE sector = 'Industrials';
UPDATE sector SET color = '#fd7038' WHERE sector = 'Materials';
UPDATE sector SET color = '#fda12f' WHERE sector = 'Consumer Staples';

UPDATE sector SET color = '#fdd214' WHERE sector = 'Energy';

UPDATE sector SET color = '#f3fd1e' WHERE sector = 'Health Care';
UPDATE sector SET color = '#52cc25' WHERE sector = 'Communication Services';
UPDATE sector SET color = '#00a537' WHERE sector = 'Consumer Discretionary';
UPDATE sector SET color = '#00a89b' WHERE sector = 'Utilities';
UPDATE sector SET color = '#2756c1' WHERE sector = 'Real Estate';

UPDATE sector SET color = '#33009c' WHERE sector = 'Information Technology';
UPDATE sector SET color = '#6d0088' WHERE sector = 'Financials';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE sector DROP COLUMN color;
