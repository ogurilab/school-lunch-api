DROP INDEX `idx_users_city_code` ON `users`;

DROP INDEX `idx_users_username` ON `users`;

DROP INDEX `idx_external_data_sources_year_status` ON `external_data_sources`;

DROP INDEX `idx_external_data_sources_city_code_status` ON `external_data_sources`;

DROP INDEX `idx_external_data_sources_city_code` ON `external_data_sources`;

DROP INDEX `idx_external_data_sources_dataset_id` ON `external_data_sources`;

DROP INDEX `idx_cities_city_name` ON `cities`;

DROP INDEX `idx_dishes_name` ON `dishes`;

DROP INDEX `idx_dishes_menu_id` ON `dishes`;

DROP INDEX `idx_menus_offered_at` ON `menus`;

DROP INDEX `idx_menus_city_code_offered_at` ON `menus`;

DROP TABLE `users`;

DROP TABLE `external_data_sources`;

DROP TABLE `cities`;

DROP TABLE `dishes_allergens`;

DROP TABLE `allergens`;

DROP TABLE `dishes`;

DROP TABLE `menus`;