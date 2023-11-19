CREATE TABLE `regions` (
  `id` SMALLINT PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `menus` (
  `id` varchar(255) PRIMARY KEY,
  `offered_at` DATE NOT NULL COMMENT '給食の提供日',
  `region_id` SMALLINT NOT NULL,
  `photo_url` varchar(255),
  `wikimedia_commons_url` varchar(255),
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `elementary_school_calories` int,
  `junior_high_school_calories` int
);

CREATE TABLE `dishes` (
  `id` varchar(255) PRIMARY KEY,
  `menu_id` varchar(255) NOT NULL,
  `name` varchar(255) UNIQUE,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `allergens` (
  `id` SMALLINT PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) UNIQUE NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `dishes_allergens` (
  `allergen_id` SMALLINT NOT NULL,
  `dishe_id` varchar(255) NOT NULL,
  PRIMARY KEY (`dishe_id`, `allergen_id`)
);

CREATE INDEX `menus_index_0` ON `menus` (`region_id`);

CREATE UNIQUE INDEX `menus_index_1` ON `menus` (`region_id`, `offered_at`);

CREATE INDEX `dishes_index_2` ON `dishes` (`menu_id`);

CREATE INDEX `dishes_index_3` ON `dishes` (`name`);