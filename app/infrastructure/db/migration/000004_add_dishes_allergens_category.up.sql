ALTER TABLE `dishes_allergens`
ADD COLUMN `category` TINYINT NOT NULL;

ALTER TABLE `dishes_allergens` DROP PRIMARY KEY;

ALTER TABLE `dishes_allergens`
ADD PRIMARY KEY (`dish_id`, `allergen_id`, `category`);