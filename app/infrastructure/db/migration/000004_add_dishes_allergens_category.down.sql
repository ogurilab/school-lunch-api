ALTER TABLE `dishes_allergens` DROP PRIMARY KEY;

ALTER TABLE `dishes_allergens` DROP COLUMN `category`;

ALTER TABLE `dishes_allergens`
ADD PRIMARY KEY (`dish_id`, `allergen_id`);