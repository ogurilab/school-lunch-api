ALTER TABLE menus DROP CONSTRAINT `menus_city_code_offers_at_unique`;

ALTER TABLE menus DROP COLUMN `city_code`;

DROP TABLE cities;