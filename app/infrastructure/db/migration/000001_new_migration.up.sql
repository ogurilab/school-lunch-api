CREATE TABLE `menus` (
  `id` varchar(255) PRIMARY KEY,
  `offered_at` DATE NOT NULL COMMENT '給食の提供日',
  `photo_url` varchar(255),
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `elementary_school_calories` int NOT NULL DEFAULT 0,
  `junior_high_school_calories` int NOT NULL DEFAULT 0,
  `city_code` SMALLINT NOT NULL
);

CREATE TABLE `dishes` (
  `id` varchar(255) PRIMARY KEY,
  `menu_id` varchar(255) NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `allergens` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) UNIQUE NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `dishes_allergens` (
  `allergen_id` INT NOT NULL,
  `dish_id` varchar(255) NOT NULL,
  PRIMARY KEY (`dish_id`, `allergen_id`)
);

CREATE TABLE `cities` (
  `city_code` SMALLINT PRIMARY KEY,
  `city_name` VARCHAR(100) NOT NULL,
  `prefecture_code` SMALLINT NOT NULL,
  `prefecture_name` VARCHAR(100) NOT NULL,
  `school_lunch_info_available` boolean NOT NULL DEFAULT FALSE COMMENT '給食のデータが登録されているかどうか'
);

CREATE TABLE `external_data_sources` (
  `source_id` INT PRIMARY KEY AUTO_INCREMENT,
  `city_code` INT NOT NULL,
  `dataset_id` VARCHAR(255) UNIQUE NOT NULL COMMENT 'linkedDataのdatasetのURL',
  `year` INT NOT NULL,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` VARCHAR(50) NOT NULL DEFAULT "Inactive" COMMENT 'Status of the data source: Active (currently in use), Inactive (not in use), Updating (currently being updated), Error (an error has occurred)',
  `category` varchar(50) NOT NULL DEFAULT "menu" COMMENT 'menu or dish or allergens',
  `description` TEXT
);

CREATE TABLE `users` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `username` VARCHAR(50) UNIQUE NOT NULL,
  `hashed_password` VARCHAR(255) NOT NULL,
  `email` VARCHAR(100) UNIQUE NOT NULL,
  `role` VARCHAR(50) NOT NULL DEFAULT "guest" COMMENT 'User roles: municipality (for municipal staff), admin (for system administrators), guest (for general users)',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `city_code` SMALLINT NOT NULL DEFAULT 0
);

CREATE INDEX `idx_menus_offered_at` ON `menus` (`offered_at`);

CREATE UNIQUE INDEX `idx_menus_city_code_offered_at` ON `menus` (`city_code`, `offered_at`);

CREATE INDEX `idx_dishes_menu_id` ON `dishes` (`menu_id`);

CREATE INDEX `idx_dishes_name` ON `dishes` (`name`);

CREATE INDEX `idx_cities_city_name` ON `cities` (`city_name`);

CREATE UNIQUE INDEX `idx_external_data_sources_dataset_id` ON `external_data_sources` (`dataset_id`);

CREATE INDEX `idx_external_data_sources_city_code` ON `external_data_sources` (`city_code`);

CREATE INDEX `idx_external_data_sources_city_code_status` ON `external_data_sources` (`city_code`, `status`);

CREATE INDEX `idx_external_data_sources_year_status` ON `external_data_sources` (`year`, `status`);

CREATE UNIQUE INDEX `idx_users_username` ON `users` (`username`);

CREATE INDEX `idx_users_city_code` ON `users` (`city_code`);