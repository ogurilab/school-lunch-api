CREATE TABLE menu_dishes (
  menu_id varchar(255) NOT NULL,
  dish_id varchar(255) NOT NULL,
  PRIMARY KEY (menu_id, dish_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

ALTER TABLE dishes DROP COLUMN menu_id;