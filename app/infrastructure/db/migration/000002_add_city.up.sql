CREATE TABLE cities (
  city_code INT PRIMARY KEY,
  city_name VARCHAR(100) NOT NULL,
  prefecture_code INT NOT NULL,
  prefecture_name VARCHAR(100) NOT NULL,
  school_lunch_info_available boolean NOT NULL DEFAULT FALSE
);

ALTER TABLE menus
ADD COLUMN `city_code` INT NOT NULL DEFAULT 0;

ALTER TABLE menus
ADD CONSTRAINT `menus_city_code_offers_at_unique` UNIQUE (`city_code`, `offered_at`);